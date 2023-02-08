FROM golang:1.19.4-alpine3.17 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . /src

RUN CGO_ENABLED=0 go build -o ./bin/gleragames ./cmd/app

FROM scratch

WORKDIR /app

COPY --from=build /src/bin/gleragames /app/gleragames
COPY --from=build /src/config /app/config/

ENTRYPOINT ["/app/gleragames"]
