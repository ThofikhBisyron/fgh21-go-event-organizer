FROM golang:1.23.0

WORKDIR /app 

COPY . /app/

RUN go mod tidy

RUN make migrate:reset

RUN go build -o backend

EXPOSE 8080

ENTRYPOINT ["/app/backend"]