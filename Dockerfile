FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM golang as dev

RUN mkdir app

WORKDIR /app

COPY --from=builder ./app .

RUN go get github.com/pilu/fresh

CMD [ "fresh" ]