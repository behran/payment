FROM golang:1.21.1

WORKDIR /go/src/app/

COPY src/ .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/payment/main.go && chmod +x app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o migration cmd/cli/main.go && chmod +x cli

FROM alpine:3.19.0

WORKDIR /usr/local/bin

COPY --from=builder /go/src/app/app app
COPY --from=builder /go/src/app/cli cli

ENTRYPOINT ["/usr/local/bin/app"]

EXPOSE 80
