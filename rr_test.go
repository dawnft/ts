package rr

import (
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	ts := TS{
		Resolution: 10 * time.Second,
		Size:       60,
		Buckets:    make(map[int64]*Bucket, 0),
	}

	now := time.Now()

	ts.Insert(now, 100)
	if *ts.Get(now).V != 100 {
		t.Fatal("data was not stored")
	}

	buckets, _ := ts.GetRange(now, now.Add(4*time.Hour))
	if buckets == nil {
		t.Fatal("no buckets found")
	}
}
