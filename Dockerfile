FROM golang:1.21

WORKDIR /user/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app

CMD ["app"]