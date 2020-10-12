package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"website_status_checker/database"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

type responseType struct {
	ID               uuid.UUID     `json:"id"`
	URLLink          string        `json:"url"`
	CrawlTimeout     time.Duration `json:"crawl_timeout"`
	Frequency        int           `json:"frequency"`
	FailureThreshold int           `json:"failure_threshold"`
	Status           string        `json:"status"`
	FailureCount     int           `json:"failure_count"`
}

func TestGetUrl(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockReposController(ctrl)
	setRepoController(mockRepo)
	s := "1343da6e-525d-4755-8f50-5984da54c3f0"
	id, _ := uuid.FromString(s)
	b := database.Pingdom{
		ID:               id,
		URLLink:          "http://stackoverflow.com",
		CrawlTimeout:     30,
		FailureThreshold: 20,
		Frequency:        10,
		Status:           "active",
		FailureCount:     0,
	}
	mockRepo.EXPECT().DatabaseGet(id).Return(b, nil)

	r := gin.Default()
	r.GET("/urls/:id", GetUrl)

	req, err := http.NewRequest("GET", "http://localhost:8080/urls/1343da6e-525d-4755-8f50-5984da54c3f0", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, b.URLLink, response.URLLink)
	assert.Equal(t, b.Status, response.Status)
	assert.Equal(t, b.Frequency, response.Frequency)
	assert.Equal(t, b.CrawlTimeout, response.CrawlTimeout)
	assert.Equal(t, b.FailureCount, response.FailureCount)
	assert.Equal(t, b.FailureThreshold, response.FailureThreshold)
}
func TestCreateUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockReposController(ctrl)
	setRepoController(mockRepo)
	s := "1343da6e-525d-4755-8f50-5984da54c3f0"
	id, _ := uuid.FromString(s)
	var dt time.Duration = 30 * time.Nanosecond
	b := database.Pingdom{
		ID:               id,
		URLLink:          "http://stackoverflow.com",
		CrawlTimeout:     dt,
		FailureThreshold: 20,
		Frequency:        10,
		Status:           "active",
	}
	mockRepo.EXPECT().DatabaseCreate("http://stackoverflow.com", dt, 10, 20).Return(b, nil)

	r := gin.Default()
	r.POST("/urls", CreateUrl)

	req, err := http.NewRequest("POST", "http://localhost:8080/urls", strings.NewReader(`{
		"url": "http://stackoverflow.com",
		"crawl_timeout": 30,
		"frequency": 10,
		"failure_threshold": 20
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, b.URLLink, response.URLLink)
	assert.Equal(t, b.CrawlTimeout, response.CrawlTimeout)
	assert.Equal(t, b.Frequency, response.Frequency)
}
func TestDeleteUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockReposController(ctrl)
	setRepoController(mockRepo)
	s := "1343da6e-525d-4755-8f50-5984da54c3f0"
	id, _ := uuid.FromString(s)

	mockRepo.EXPECT().DatabaseDelete(id).Return(nil)

	r := gin.Default()
	r.DELETE("urls/:id", Deleteurl)

	req, err := http.NewRequest("DELETE", "http://localhost:8080/urls/1343da6e-525d-4755-8f50-5984da54c3f0", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestUpdateUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockReposController(ctrl)
	setRepoController(mockRepo)
	s := "1343da6e-525d-4755-8f50-5984da54c3f0"
	id, _ := uuid.FromString(s)
	var dt time.Duration = 30 * time.Nanosecond
	b := database.Pingdom{
		ID:               id,
		URLLink:          "http://stackoverflow.com",
		CrawlTimeout:     dt,
		FailureThreshold: 20,
		Frequency:        10,
		Status:           "active",
	}
	mockRepo.EXPECT().DatabaseUpdate(id, dt, 10, 20).Return(b, nil)

	r := gin.Default()
	r.PATCH("urls/:id", Updateurl)

	req, err := http.NewRequest("PATCH", "http://localhost:8080/urls/1343da6e-525d-4755-8f50-5984da54c3f0", strings.NewReader(`{
		"crawl_timeout": 30,
		"frequency": 10,
		"failure_threshold": 20
	  }`))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response responseType
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, b.ID, response.ID)
	assert.Equal(t, b.URLLink, response.URLLink)
	assert.Equal(t, b.Status, response.Status)
	assert.Equal(t, b.Frequency, response.Frequency)
	assert.Equal(t, b.CrawlTimeout, response.CrawlTimeout)
	assert.Equal(t, b.FailureCount, response.FailureCount)
	assert.Equal(t, b.FailureThreshold, response.FailureThreshold)
}
