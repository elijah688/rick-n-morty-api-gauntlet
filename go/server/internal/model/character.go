package model

import (
	"errors"
	"riki/internal/utils"
	"strconv"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type Character struct {
	ID       *int       `json:"id,omitempty"`
	Name     *string    `json:"name,omitempty"`
	Status   *string    `json:"status,omitempty"`
	Species  *string    `json:"species,omitempty"`
	Type     *string    `json:"type,omitempty"`
	Gender   *string    `json:"gender,omitempty"`
	Origin   *Location  `json:"origin,omitempty"`
	Location *Location  `json:"location,omitempty"`
	Image    *string    `json:"image,omitempty"`
	Episodes *[]int     `json:"episodes,omitempty"`
	URL      *string    `json:"url,omitempty"`
	Created  *time.Time `json:"created,omitempty"`
}

func (c *Character) ToRecord() goqu.Record {
	r := goqu.Record{
		"name":        c.Name,
		"status":      c.Status,
		"species":     c.Species,
		"type":        c.Type,
		"gender":      c.Gender,
		"origin_id":   0,
		"location_id": 0,
		"image":       c.Image,
		"url":         c.URL,
		"created":     c.Created,
	}

	if lID := getLocationID(c.Location); lID != nil {
		r["location_id"] = *lID

	}

	if oID := getLocationID(c.Origin); oID != nil {
		r["origin_id"] = *oID
	}

	if c.ID != nil {
		r["id"] = c.ID
	}

	return r
}

func getLocationID(l *Location) *int {

	if l == nil {
		return nil
	}

	if l.ID != nil {
		return l.ID
	}

	if l.URL != nil {

		basename := utils.Basename(*l.URL)
		if locationID, err := strconv.Atoi(basename); err == nil {
			return &locationID
		}
	}

	return nil

}

type Characters []Character

func (cs Characters) ToRecords() []goqu.Record {
	records := make([]goqu.Record, len(cs))
	for i, c := range cs {
		records[i] = c.ToRecord()
	}

	return records
}

func (c *Character) Validate() error {
	if c.Name == nil || *c.Name == "" {
		return errors.New("name is required")
	}

	return nil

}
