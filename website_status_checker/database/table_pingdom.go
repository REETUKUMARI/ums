package database

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Pingdom struct {
	ID               uuid.UUID     `json:"id" gorm:"primary_key"`
	URLLink          string        `json:"url"`
	CrawlTimeout     time.Duration `json:"crawl_timeout"`
	Frequency        int           `json:"frequency"`
	FailureThreshold int           `json:"failure_threshold"`
	Status           string        `json:"status"`
	FailureCount     int           `json:"failure_count"`
}

func (base *Pingdom) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("ID", uuid)
}
