# OverlappingDates

### Local Testing

```
go run main.go
```

Use curl or Postman to send a POST request against http://localhost:8081/date-range/overlap
#### Example Request
```
{
    "range1": {
        "start_date": "2022-01-01T00:00:00Z",
        "end_date": "2022-02-01T00:00:00Z"
    },
    "range2": {
        "start_date": "2022-01-15T00:00:00Z",
        "end_date": "2022-02-15T00:00:00Z"
    }
}
```
#### Example Response
```
{
    "overlap": true
}
```

### Running unit test

```
go test ./pkg/...
```

### Running integration test
```
go test ./tests/...
```

### Building Proto files

```
protoc --go_out=. --go_opt=paths=source_relative api.proto
protoc --go_out=. --go_opt=paths=source_relative json_options.proto
``` 

### Docker Container

#### Building Docker Image
```docker build -t overlapping-date-service .```

#### Running Docker Image
```docker run -p 8080:8080 overlapping-date-service``` 