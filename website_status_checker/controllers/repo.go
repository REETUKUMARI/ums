package controllers

import (
	"time"
	"website_status_checker/database"

	uuid "github.com/satori/go.uuid"
)

type ReposController interface {
	DatabaseGet(id uuid.UUID) (database.Pingdom, error)
	DatabaseGets(url *[]database.Pingdom) error
	DatabaseCreate(s string, c time.Duration, fre int, fail int) (database.Pingdom, error)
	DatabaseDelete(id uuid.UUID) error
	DatabaseUpdate(id uuid.UUID, c time.Duration, fre int, fail int) (database.Pingdom, error)
}

type MonitorRepo struct{}

var repo ReposController

func setRepoController(repoType ReposController) {
	repo = repoType
}
func init() {
	setRepoController(&MonitorRepo{})
}

func (rp *MonitorRepo) DatabaseGet(id uuid.UUID) (database.Pingdom, error) {
	var url database.Pingdom
	result := database.DB.Where("id = ?", id).First(&url)
	if result.Error != nil {
		return database.Pingdom{}, result.Error
	}
	return url, nil
}
func (rp *MonitorRepo) DatabaseGets(url *[]database.Pingdom) error {
	return database.DB.First(&url).Error
}
func (rp *MonitorRepo) DatabaseCreate(s string, c time.Duration, fre int, fail int) (database.Pingdom, error) {
	url := database.Pingdom{URLLink: s, CrawlTimeout: c, Frequency: fre, FailureThreshold: fail}
	url.Status = checkurl(url.URLLink, url.CrawlTimeout)

	err := database.DB.Create(&url).Error
	return url, err
}
func (rp *MonitorRepo) DatabaseDelete(id uuid.UUID) error {
	var url database.Pingdom
	d := database.DB.Where("id = ?", id).Delete(&url)
	return d.Error
}
func (rp *MonitorRepo) DatabaseUpdate(id uuid.UUID, c time.Duration, fre int, fail int) (database.Pingdom, error) {
	var url database.Pingdom
	result := database.DB.Where("id = ?", id).First(&url)
	if result.Error != nil {
		return database.Pingdom{}, result.Error
	}
	url.CrawlTimeout = c
	url.Frequency = fre
	url.FailureThreshold = fail
	database.DB.Save(&url)
	return url, nil
}
