FROM golang:1.17

RUN apt update; \
    apt install make

RUN mkdir form3-client-go

WORKDIR /form3-client-go
COPY . .

RUN CGO_ENABLED=0 go get ./...
RUN go test ./f3client/... -run=^Test_Unit_ -cover -v
RUN go test ./f3client/... -p 1 -run=^Test_Integration_ -v -cover



