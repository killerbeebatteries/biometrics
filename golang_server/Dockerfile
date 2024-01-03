# based off https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/
FROM golang:1.20 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
CMD ["air"]

FROM base as built

WORKDIR /go/app/api
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go build -o /tmp/api-server ./*.go

FROM alpine:latest as prod

WORKDIR /app
COPY --from=built /go/app/api/*.html ./
COPY --from=built /tmp/api-server /usr/bin/api-server
CMD ["/usr/bin/api-server"]
