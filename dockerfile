FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o backend

FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/backend /app/backend

EXPOSE 8080

ENTRYPOINT ["/app/backend"]


# Normal
# FROM golang:1.23.0

# WORKDIR /app 

# COPY . /app/

# RUN go mod tidy

# RUN go build -o backend

# EXPOSE 8080

# ENTRYPOINT ["/app/backend"]