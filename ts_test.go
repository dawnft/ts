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
	buckets = s.FromDuration(10 * time.Second)
	if len(buckets) != 1 {
		t.Fatalf("we should have 1 bucket but we have %d\n", len(buckets))
	}

	if buckets[0].V != 100 {
		t.Fatal("data was not stored")
	}

	if buckets[0].C != 1 {
		t.Fatal("counter was not updated")
	}

	s.Insert(now, 100)
	if buckets[0].V != 200 {
		t.Fatal("data was not updated on second insert")
	}

	if buckets[0].C != 2 {
		t.Fatal("counter was not updated on second insert")
	}

	buckets = s.FromDuration(30 * time.Minute)
	if len(buckets) != 180 {
		t.Fatalf("buckets = %d, not 180", len(buckets))
	}
}
