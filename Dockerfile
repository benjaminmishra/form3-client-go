FROM golang:1.17

RUN apt update; \
    apt install make

RUN mkdir form3-client-go

WORKDIR /form3-client-go
COPY . .

RUN CGO_ENABLED=0 go get ./...
CMD ["make","test"]



