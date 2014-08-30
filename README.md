ts
==

lightweight time-series lib for go

## Installation

```
$ go get github.com/gorsuch/ts
```

## Usage

```go
// create a series 4h long, with 5s buckets
ts := ts.NewTS(4 * time.Hour, 5 * time.Second)

// insert some data
ts.Insert(time.Now(), 42)

// fetch the last hour of data
now := time.Now()
anHourAgo := now.Add(-1 * time.Hour)
buckets := ts.Range(anHourAgo, now)
```
