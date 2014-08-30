package rr

import (
	"fmt"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	ts := TS{
		Resolution: 10,
		Size:       60,
		Buckets:    make(map[int64]*Bucket, 0),
	}

	now := time.Now()
	epoch := now.Unix()

	ts.Insert(epoch, 100)
	if *ts.Get(epoch).Value != 100 {
		t.Fatal("data was not stored")
	}

	buckets, _ := ts.GetRange(epoch, now.Add(4*time.Hour).Unix())
	if buckets == nil {
		t.Fatal("no buckets found")
	}
	fmt.Println(len(buckets))
}
