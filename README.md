ts [![Build Status](https://travis-ci.org/gorsuch/ts.svg)](https://travis-ci.org/gorsuch/ts)
==

lightweight time-series lib for go

## Goals

* metric insert should be idempotent
* * easy to understand
* no persistennce, that's up to you
* no metric naming schemes or discoverability, that's up to you

## Installation

```
$ go get github.com/gorsuch/ts
```

## Usage

```go
// create a series 4h long, with 5s buckets
s := ts.NewSeries(4 * time.Hour, 5 * time.Second)

// insert some data
s.Insert(time.Now(), 42)

// fetch the last hour of data
now := time.Now()
anHourAgo := now.Add(-1 * time.Hour)
buckets := s.Range(anHourAgo, now)
```

## TODO

* [ ] buid a higher abstraction to rollup one series into another
