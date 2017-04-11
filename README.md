[![Build Status](https://travis-ci.org/cagiti/go-tawerin.svg?branch=master)](https://travis-ci.org/cagiti/go-tawerin)

# go-tawerin
Tawerin website using templating, gorilla/mux router and newrelic monitoring

## Build the docker container
```
docker build -t go-tawerin .
```

## Run the docker container
```
docker run -d --name go-tawerin -p 80:8080 -e NEWRELIC_LICENSE_KEY="<YOUR_LICENSE_KEY_HERE>" -e NEWRELIC_APP_NAME="go-tawerin" -e PORT="8080" go-tawerin
```

## Run go app
```
godep save
go install
go-tawerin
```
