FROM golang:1.17-alpine

WORKDIR /go/src/game-currency

COPY . .

RUN CGO_ENABLED=0 go build

ENTRYPOINT [ "./game-currency" ]
