FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

#stage 2

FROM alpine:latest

WORKDIR /app

COPY --from=build /app .

EXPOSE 3000

CMD ["./main"]