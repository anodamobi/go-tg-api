FROM golang:1.12.1-stretch

WORKDIR $GOPATH/src/github.com/anodamobi/glance-backend/

COPY . .

RUN go build -o glance -v ./cmd/main.go