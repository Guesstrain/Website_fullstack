FROM golang:1.21.5

RUN mkdir /GoWeb

ADD . /GoWeb

WORKDIR /GoWeb

COPY go.* ./

RUN go mod download && go mod verify

RUN go build -o app .

EXPOSE 8080