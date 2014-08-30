package ts

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	ts := NewTS(1*time.Hour, 10*time.Second)

	now := time.Now()

	buckets := ts.Range(now, now)
	if buckets == nil {
		t.Fatal("no buckets found")
	}

	ts.Insert(now, 100)
	buckets = ts.Range(now, now)
	if len(buckets) != 1 {
		t.Fatalf("we should have 1 bucket but we have %d\n", len(buckets))
	}

	if *buckets[0].V != 100 {
		t.Fatal("data was not stored")
	}
}
