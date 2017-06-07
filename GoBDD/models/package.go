package models

import "time"

// Package : here you tell us what Salutation is
type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	UpdatedAt     time.Time
	LastUpdatedBy string
	ReadMe        string
	Tags          []string
	Categories    []string
}
