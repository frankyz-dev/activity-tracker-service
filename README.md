# activity-tracker-service
Activity Tracker Service holds the APIs and database integrations to support a very basic set of activity tracking features.
There is another package activity-tracking-react that holds the react web implementation which makes calls to this service.

This package is built using golang, of which I have just recently learned.

## Data Model

## Features & APIs

## Run it yourself
This expects a postgres database

windows: 
```
cd activity-tracker-service/cmd/server
go mod tidy
go build
$env:DB_HOST = 'localhost'
$env:DB_PORT = '5432'
$env:DB_USER = ''
$env:DB_PASSWORD = ''
$env:DB_NAME = ''
./server.exe
```