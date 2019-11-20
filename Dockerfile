FROM golang:1.12.1-stretch

WORKDIR $GOPATH/src/github.com/anodamobi/go-tg-api/

COPY . .

RUN go build -o go-tg-api -v ./cmd/main.go