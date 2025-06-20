package server

import (
	"fmt"
	"sync"
)

type Record struct {
	Value []byte `json:"value"`
	Off   uint64 `json:"offset"`
}

type Log struct {
    mu sync.Mutex
    records []Record
}

func NewLog() *Log {
    return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
    c.mu.Lock()
    defer c.mu.Unlock()

    record.Off = uint64(len(c.records))
    c.records = append(c.records, record)
    return record.Off, nil
}

func (c *Log) Read(off uint64) (Record, error) {
    c.mu.Lock()
    defer c.mu.Unlock()

    if off >= uint64(len(c.records)) {
        return Record{}, fmt.Errorf("offset out of range")
    }
    return c.records[off], nil
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")