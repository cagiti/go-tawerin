FROM golang:1.14 as builder

WORKDIR /usr/local/go/src/github.com/cagiti/go-tawerin

COPY go.mod .
COPY go.sum .

ENV GO111MODULE=on

RUN go mod download

COPY web/ web
COPY pkg/ pkg

RUN env CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o go-tawerin-linux-amd64 web/tawerin/tawerin.go

FROM busybox:glibc as production

ENV XDG_DATA_HOME=/home/.local/share
ENV PATH=$XDG_DATA_HOME/tawerin:$PATH

COPY --from=builder /usr/local/go/src/github.com/cagiti/go-tawerin/go-tawerin-linux-amd64 $XDG_DATA_HOME/tawerin/go-tawerin
COPY static/ $XDG_DATA_HOME/tawerin/static
COPY templates/ $XDG_DATA_HOME/tawerin/templates

CMD ["go-tawerin"]
