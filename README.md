# Benchmark RethinkDB with GO client driver

NOTE: this is the very first project that I use Docker

## How to run benchmark

### Build docker image for bm-rethinkdb
At root of source code, run:

    docker build -t bm-rethinkdb .

### Run benchmark
Firstly start RethinkDB

    docker run --name rethinkdb -d rethinkdb

At root of source code, run test by following command

    docker run --rm --link rethinkdb:rethinkdb_alias -v "$PWD:/go/src/bm-rethinkdb" bm-rethinkdb
