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

	b := buckets[0]

	if b.Value != 100 {
		t.Fatal("data was not stored")
	}

	if b.Count != 1 {
		t.Fatal("counter was not updated")
	}

	s.Insert(now, 101)
	if b.Value != 201 {
		t.Fatal("data was not updated on second insert")
	}

	if b.Count != 2 {
		t.Fatal("counter was not updated on second insert")
	}

	if b.Min != 100 {
		t.Fatal("min not set to what it should be")
	}

	if b.Max != 101 {
		t.Fatal("max not set to what it should be")
	}

	buckets = s.FromDuration(30 * time.Minute)
	if len(buckets) != 180 {
		t.Fatalf("buckets = %d, not 180", len(buckets))
	}
}
