FROM golang:1.18.1-alpine
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./server cmd/main.go

FROM alpine
COPY --from=0 /src/server /usr/bin/server
CMD ["server"]
