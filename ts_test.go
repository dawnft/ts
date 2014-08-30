package ts

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	s := NewSeries(1*time.Hour, 10*time.Second)

	now := time.Now()

	buckets := s.Range(now, now)
	if buckets == nil {
		t.Fatal("no buckets found")
	}

	s.Insert(now, 100)
	buckets = s.Range(now, now)
	if len(buckets) != 1 {
		t.Fatalf("we should have 1 bucket but we have %d\n", len(buckets))
	}

	if *buckets[0].V != 100 {
		t.Fatal("data was not stored")
	}
}
