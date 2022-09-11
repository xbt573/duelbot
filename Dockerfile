FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build

FROM scratch AS base

WORKDIR /
COPY --from=build /app/duelbot /duelbot

CMD /duelbot
