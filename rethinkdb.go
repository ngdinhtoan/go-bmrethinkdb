package bmrethinkdb

import (
	"flag"
	"fmt"
	"log"

	r "github.com/dancannon/gorethink"
)

var (
	host string     // RethinkDB host
	port int        // RethinkDB client port
	conn *r.Session // connection to RethinkDB

	dbName    = "test"    // RethinkDB database name
	tableName = "dbevent" // RethinkDB table name
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "RethinkDB host")
	flag.IntVar(&port, "port", 28015, "RethinkDB client port")
	flag.Parse()

	var err error
	opts := r.ConnectOpts{
		Address:  fmt.Sprintf("%s:%d", host, port),
		Database: dbName,
		MaxIdle:  10,
		MaxOpen:  10,
	}

	if conn, err = r.Connect(opts); err != nil {
		log.Fatalf("could not connect to RethinkDB, error: %v", err)
	}

	// pre-create table for test,
	// if table already exists, just ignore error message :-)
	r.DB(dbName).TableCreate(tableName).RunWrite(conn)
}

func write(data map[string]interface{}, iopt r.InsertOpts) error {
	term := r.Table(tableName).Insert(data, iopt)
	// _, err := term.RunWrite(conn)
	err := term.Exec(conn, r.ExecOpts{NoReply: true})
	return err
}

// Write data to RethinkDB
func Write(data map[string]interface{}) error {
	return write(data, r.InsertOpts{})
}

// SoftWrite data to RethinkDB in with soft durability
func SoftWrite(data map[string]interface{}) error {
	return write(data, r.InsertOpts{Durability: "soft"})
}
