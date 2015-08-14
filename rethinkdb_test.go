package bmrethinkdb

import (
	"sync"
	"testing"
)

func TestWrite(t *testing.T) {
	data := map[string]interface{}{
		"event_id":   1,
		"event_name": "test",
	}

	if err := Write(data); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkWrite(b *testing.B) {
	for i := 0; i < b.N; i = i + 1 {
		data := map[string]interface{}{
			"event_id":   i,
			"event_name": "test",
		}

		if err := Write(data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSoftWrite(b *testing.B) {
	for i := 0; i < b.N; i = i + 1 {
		data := map[string]interface{}{
			"event_id":   i,
			"event_name": "test",
		}

		if err := SoftWrite(data); err != nil {
			b.Fatal(err)
		}
	}
}

func runParallelWrite(b *testing.B, fn func(data map[string]interface{}) error) {
	var (
		mu sync.Mutex
		i  int
	)

	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			i++
			mu.Unlock()

			data := map[string]interface{}{
				"event_id":   i,
				"event_name": "test",
			}

			if err := fn(data); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkWriteParallel(b *testing.B) {
	runParallelWrite(b, Write)
}

func BenchmarkSoftWriteParallel(b *testing.B) {
	runParallelWrite(b, SoftWrite)
}
