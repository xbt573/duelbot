FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build

FROM alpine:latest

RUN apk add libc6-compat

WORKDIR /
COPY --from=build /app/duelbot .

CMD /duelbot
