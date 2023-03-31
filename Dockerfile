# syntax=docker/dockerfile:1

FROM golang:1.19.7-alpine3.17

WORKDIR /app

COPY go.mod  go.sum ./
RUN go mod download

COPY ./cmd/PracticeProject/*.go ./cmd/PracticeProject/
COPY ./pkg/ ./pkg/

RUN go build ./cmd/PracticeProject/

EXPOSE 8081

# CMD [ "main" ]
CMD [ "/app/PracticeProject" ]