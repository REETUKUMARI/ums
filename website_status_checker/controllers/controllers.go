package controllers

import (
	"fmt"
	"net/http"
	"time"
	"website_status_checker/database"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CreateUrlInput struct {
	URLLink          string        `json:"url" binding:"required"`
	CrawlTimeout     time.Duration `json:"crawl_timeout" binding:"required"`
	Frequency        int           `json:"frequency" binding:"required"`
	FailureThreshold int           `json:"failure_threshold" binding:"required"`
}

type UpdateUrlInput struct {
	CrawlTimeout     time.Duration `json:"crawl_timeout"`
	Frequency        int           `json:"frequency"`
	FailureThreshold int           `json:"failure_threshold"`
}

func GetUrls(c *gin.Context) {
	var url []database.Pingdom
	if err := repo.DatabaseGets(&url); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		database.DB.Find(&url)
		c.JSON(http.StatusOK, url)
	}
}
func GetUrl(c *gin.Context) {
	id, err := uuid.FromString(c.Params.ByName("id"))
	res, err := repo.DatabaseGet(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func CreateUrl(c *gin.Context) {
	var urll CreateUrlInput
	c.BindJSON(&urll)
	url, err := repo.DatabaseCreate(urll.URLLink, urll.CrawlTimeout, urll.Frequency, urll.FailureThreshold)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		//models.DB.Create(&url)
		c.JSON(http.StatusOK, url)
	}
}
func Updateurl(c *gin.Context) {
	id := StringToUUID(c.Params.ByName("id"))
	var input database.Pingdom
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url, err := repo.DatabaseUpdate(id, input.CrawlTimeout, input.Frequency, input.FailureThreshold)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, url)
}
func Deleteurl(c *gin.Context) {
	idd := c.Params.ByName("id")
	id := StringToUUID(idd)
	d := repo.DatabaseDelete(id)
	if d != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	//c.JSON(200, http.StatusNoContent)
	c.Writer.WriteHeader(http.StatusNoContent)
}

func Checklink() {
	rows, err := database.DB.Raw("select * from pingdoms WHERE status != ?", "inactive").Rows()
	if err != nil {
		fmt.Println("connection failed")
	}

	defer rows.Close()

	var (
		id              string
		urllink         string
		crawltimeout    time.Duration
		frequency       int
		failurethresold int
		status          string
		failurecount    int
	)

	c := make(chan string)
	for rows.Next() {
		rows.Scan(&id, &urllink, &crawltimeout, &frequency, &failurethresold, &status, &failurecount)
		//fmt.Println(id, urllink, crawltimeout, frequency, failurethresold, status, failurecount)
		go checkLink(urllink, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(1 * time.Second)
			if link != "" {
				checkLink(link, c)
			}
		}(l)
	}
}

func checkurl(url string, crawltimeout time.Duration) string {
	client := http.Client{
		Timeout: crawltimeout * time.Second,
	}
	_, err := client.Get(url)
	if err != nil {
		return "inactive"
	} else {
		return "active"
	}
}

func checkLink(urllink string, c chan string) {
	var p database.Pingdom
	database.DB.First(&p, "url_link  = ?", urllink)
	client := http.Client{
		Timeout: p.CrawlTimeout * time.Second,
	}
	_, err := client.Get(urllink)
	if err != nil {
		failurecount := p.FailureCount + 1
		//fmt.Println(p.FailureCount)
		database.DB.Model(&p).Update("FailureCount", failurecount)
		database.DB.Model(&p).Update("Status", "inactive")
		if failurecount >= p.FailureThreshold {
			database.DB.Model(&p).Update("FailureCount", 0)
			c <- ""
		}
		//fmt.Println(urllink, "is down!")
	} else {
		if p.Status != "active" {
			database.DB.Model(&p).Update("Status", "active")
			database.DB.Model(&p).Update("FailureCount", 0)
		}
		//fmt.Println(urllink, "is up!")
	}
	c <- urllink
}
func StringToUUID(st string) uuid.UUID {
	id, err := uuid.FromString(st)
	if err != nil {
		fmt.Println("Error while Converting UID", err.Error())
	}
	return id
}
