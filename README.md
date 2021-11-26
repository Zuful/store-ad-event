# store-ad-event
A REST API allowing to save ad events and its related informations.

## Running the app
```bash
# development
$ go run *.go
```

## Test
```bash
# unit tests
$ go test
```

## Query the endpoint
Launch a POST request on the following endpoint in order to save an event:
```bash
/event
```
The body of the POST request must look like this:  
```json
{
  "type":"visible",
  "timestamp":1637885000
}
```
Launch a GET request on the following endpoint in order a list of event. There are 2 required parameters that are 
timestamps representing the time interval within which we want to retrieve de informations.
```bash
/event/1637884800/1637971199
```
Alternatively you can also add an optional query parameter : "type" and "os" in order to group the results.
```bash
/event/1637884800/1637971199?groupBy=type,os
```
