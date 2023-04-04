FROM golang:1.19-alpine AS build
RUN apk update && \
    apk add curl \
    git \
    bash \
    make \
    ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

# copy module files first so that they don't need to be downloaded again if no change
ENV GOPATH /go
COPY go.* ./
RUN go mod download
RUN go mod verify

# copy source files and build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o application main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash tzdata
ENV TZ Asia/Jakarta

# RUN mkdir -p /var/log/app
WORKDIR /app/
COPY --from=build /app/application .
COPY .env.development /app/env.development
#COPY --from=build /app/scripts/migrations/mysql ./scripts/migrations/mysql/
#COPY --from=build /app/assets ./assets/

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.6.0/wait /wait
RUN chmod +x /wait

RUN ["chmod", "+x", "./application"]
ENTRYPOINT ["./application"]