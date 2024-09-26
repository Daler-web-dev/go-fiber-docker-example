FROM golang:1.23.1


WORKDIR /user/src/trivia

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy