FROM golang:1.22 as build

WORKDIR /app
COPY . .
RUN go mod download \
    && go test ./... \
    && go build -o musiclandapi ./cmd/musicland/main.go

FROM golang:1.22

WORKDIR /app
COPY --from=build /app/musiclandapi /app/musiclandapi
