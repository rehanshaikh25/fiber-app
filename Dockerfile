# Stage 1: Build Stage
FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o main .

# Stage 2: Runtime Stage
FROM alpine:3.18

WORKDIR /app

# Copy only the built binary from the build stage
COPY --from=build /app/main .

EXPOSE 3000

CMD ["./main"]
