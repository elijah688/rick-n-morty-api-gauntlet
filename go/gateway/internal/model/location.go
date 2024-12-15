package model

import (
	"time"

	"github.com/doug-martin/goqu/v9"
)

type Location struct {
	ID        *int       `json:"id"`
	Name      *string    `json:"name"`
	URL       *string    `json:"url"`
	Type      *string    `json:"type"`
	Dimension *string    `json:"dimension"`
	Residents []string   `json:"residents"`
	Created   *time.Time `json:"created"`
}

func (l *Location) ToRecord() goqu.Record {
	return goqu.Record{
		"id":        l.ID,
		"name":      l.Name,
		"url":       l.URL,
		"type":      l.Type,
		"dimension": l.Dimension,
		"created":   l.Created,
	}
}

type Locations []Location

func (ls Locations) ToRecords() []goqu.Record {
	records := make([]goqu.Record, len(ls))
	for i, l := range ls {
		records[i] = l.ToRecord()
	}
	return records
}
