FROM golang:1.19-bullseye as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY src/ ./src

RUN go mod download
RUN go mod verify

COPY . .

RUN go build -v -o /server src/main.go


FROM golang:1.19-bullseye as runner

COPY --from=builder /server /usr/local/bin/server

ENTRYPOINT [ "server" ]
