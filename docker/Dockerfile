FROM golang:latest AS builder
ADD ./ /app
WORKDIR /app
RUN go mod tidy \
	&& CGO_ENABLED=0 go build -o ./bin/server -a -i /app/cmd/server/main.go

# final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/server /app/
ENTRYPOINT ["./server"]
