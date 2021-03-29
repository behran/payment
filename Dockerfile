FROM golang:1.16 

WORKDIR /go/src/app/

COPY src/ .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/payment/main.go && chmod +x app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o migration cmd/migration/main.go && chmod +x migration

EXPOSE 80
