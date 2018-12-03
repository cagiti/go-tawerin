[![Build Status](https://travis-ci.org/cagiti/go-tawerin.svg?branch=master)](https://travis-ci.org/cagiti/go-tawerin)

# go-tawerin
Tawerin website using templating, gorilla/mux router

## Build the docker container
```
docker build -t go-tawerin .
```

## Run the docker container
```
docker run -d \
    --name go-tawerin \
    -p 8080:8080 \
    -e PORT="8080" go-tawerin
```

## Run go app
```
make
./go-tawerin
```
