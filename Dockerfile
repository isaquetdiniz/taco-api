FROM golang:1.19-bullseye as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o server src/main.go


FROM scratch as runner

COPY --from=builder /app/server /server

ENTRYPOINT [ "/server" ]
