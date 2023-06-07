package models

import (
	"time"
)

type ListSites map[string]*SiteStatus

type SiteStatus struct {
	Status bool          `default:"false" json:"status"`
	Time   time.Duration `json:"time"`
}
