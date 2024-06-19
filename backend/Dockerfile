FROM golang:1.22-alpine
RUN apk add build-base
WORKDIR /app
COPY /backend /app/
EXPOSE 8111
ENV GO111MODULE=on
RUN go mod download
RUN go mod tidy
RUN ls -la /app/
CMD go run /app/cmd/main.go
