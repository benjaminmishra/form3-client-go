FROM golang:1.17


RUN apt update; \
    apt install make


RUN mkdir client
RUN pwd
WORKDIR /client
COPY . .

CMD ["make","test"]



