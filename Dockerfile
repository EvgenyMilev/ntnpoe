# build the server binary
FROM golang:1.15.3 AS builder
LABEL stage=server-intermediate
WORKDIR /go/src/github.com/EvgenyMilev/ntnpoe

COPY . .
RUN go build -o bin/app .

# copy the binary from builder stage; run the binary
FROM alpine:latest AS runner
WORKDIR /bin

COPY --from=builder /go/src/github.com/EvgenyMilev/ntnpoe/bin .
ENTRYPOINT ["app"]
