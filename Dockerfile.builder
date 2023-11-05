FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

RUN mkdir /app 
WORKDIR /app/ 

COPY ./go.mod .
COPY ./go.sum .
COPY ./sensorServer.go .

RUN mkdir cache
COPY ./cache/ ./cache/

RUN mkdir cronjob
COPY ./cronjob/ ./cronjob/

RUN mkdir db
COPY ./db/ ./db/

RUN mkdir logger
COPY ./logger/ ./logger/

RUN mkdir webservice
COPY ./webservice/ ./webservice/

RUN go build -o sensorServer 

EXPOSE 8080

ENTRYPOINT  /app/sensorServer