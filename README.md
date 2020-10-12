# Uptime-Monitoring-System
check whether the requested url is active or inactive

# Tech Stack Used
  - GoLang - gin (microframework)
  - MySQL
    - Gorm as orm Library
  - Docker
# Installation

### CREATE Your .env File
```
MYSQL_USER="user_name"
MYSQL_PASSWORD="password"
MYSQL_DBNAME="table_name"
MYSQL_HOST="localhost"
MYSQL_PORT="3306"
```
### 1. On Local Machine
```
git clone https://github.com/REETUKUMARI/Uptime-Monitoring-System.git
cd Uptime-Monitoring-System
```
  - Run the Following command
```
go build .
./Uptime-Monitoring-System
```
### 2. Using Docker
Run the following commands
```
git clone https://github.com/REETUKUMARI/Uptime-Monitoring-System.git
cd Uptime-Monitoring-System
```
#### Build
```
docker build . -t Uptime-Monitoring-System
```
#### Run
```
docker run -p 8080:8080 Uptime-Monitoring-System
```
### Using Docker image
Pull the image from dockerhub by executing the following command
```
docker image pull reetukumari/status-checker
```
#### Run
```
docker run -p 8080:8080 status-checker
```
# API
### Check Url:
GET /urls/:id
#### Response:
```
{
 "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”abc.com”,
  "crawl_timeout":               20,
  “frequency”:                  30, 
  “failure_threshold” :         50 
  “status”:                     “active”, 
  “failure_count”:               0

}
```
### Add Url:
POST/URL
#### Request:
```
{
  "url":                        ”abc.com”,
  "crawl_timeout":               20,
  “frequency”:                  30, 
  “failure_threshold” :         50 
}
```
#### Response:
```
{
  "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”abc.com”,
  "crawl_timeout":               20,
  “frequency”:                  30, 
  “failure_threshold” :         50 
  “status”:                     “active”, 
  “failure_count”:               0
}
```
### Update Url:
PATCH /url/:id
#### Request:
```
{
  “frequency”:                  60, 
  “status”:                     “active” 
}
```
#### Response:
```
{
  "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”abc.com”,
  "crawl_timeout":               20,
  “frequency”:                  60, 
  “failure_threshold” :         50 
  “status”:                     “active”, 
  “failure_count”:               0

}
```
### Delete URL:
DELETE /urls/:id
