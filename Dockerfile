FROM golang:1.23.0-alpine3.19 AS builder

# Install git.
RUN apk update && apk add --no-cache git tzdata
WORKDIR /app

COPY controller ./controller
COPY domain ./domain
COPY handler ./handler
COPY lib ./lib
COPY repository ./repository
COPY usecase ./usecase
COPY vendor ./vendor
COPY go.mod ./go.mod
COPY go.sum ./go.sum
COPY main.go ./main.go
COPY .env ./.env

RUN CGO_ENABLED=0 GOOS=linux DOCKER_BUILDKIT=0 go build -o xseon-zero main.go

FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
RUN apk add --no-cache qpdf
WORKDIR /app/
RUN mkdir public

COPY --from=builder app/xseon-zero .
COPY --from=builder app/.env .
EXPOSE 80
CMD ["./xseon-zero"]