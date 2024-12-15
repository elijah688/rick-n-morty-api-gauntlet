package model

import (
	"riki_gateway/internal/utils"
	"strconv"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type Episode struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	AirDate     string    `json:"air_date"`
	EpisodeCode string    `json:"episode"`
	Characters  []string  `json:"characters"`
	URL         string    `json:"url"`
	Created     time.Time `json:"created"`
}

func (e *Episode) ToRecord() goqu.Record {
	return goqu.Record{
		"id":           e.ID,
		"name":         e.Name,
		"air_date":     e.AirDate,
		"episode_code": e.EpisodeCode,
		"url":          e.URL,
		"created":      e.Created,
	}
}

type Episodes []Episode

func (es Episodes) ToRecords() []goqu.Record {
	records := make([]goqu.Record, len(es))
	for i, e := range es {
		records[i] = e.ToRecord()
	}
	return records
}

func (e *Episode) ToEpisodeCharacterListRecord() ([]goqu.Record, error) {
	records := make([]goqu.Record, len(e.Characters))
	for i, characterURL := range e.Characters {
		basename := utils.Basename(characterURL)
		characterID, err := strconv.Atoi(basename)
		if err != nil {
			return nil, err
		}
		records[i] = goqu.Record{
			"character_id": characterID,
			"episode_id":   e.ID,
		}
	}
	return records, nil
}

func (es Episodes) ToEpisodeCharacterListRecords() ([]goqu.Record, error) {
	var records []goqu.Record
	for _, e := range es {
		episodeRecords, err := e.ToEpisodeCharacterListRecord()
		if err != nil {
			return nil, err
		}
		records = append(records, episodeRecords...)
	}
	return records, nil
}
