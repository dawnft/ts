package rr

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	ts := TS{
		Resolution: 10 * time.Second,
		Duration:   1 * time.Hour,
		Buckets:    make(map[int64]*Bucket, 0),
	}

	now := time.Now()

	buckets, _ := ts.Range(now, now)
	if buckets == nil {
		t.Fatal("no buckets found")
	}

	ts.Insert(now, 100)
	buckets, _ = ts.Range(now, now)
	if len(buckets) != 1 {
		t.Fatalf("we should have 1 bucket but we have %d\n", len(buckets))
	}

	if *buckets[0].V != 100 {
		t.Fatal("data was not stored")
	}
}
