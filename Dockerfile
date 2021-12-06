FROM golang:1.15-alpine as builder

RUN apk add --no-cache git
RUN apk add --update make
RUN mkdir -p /app

# Move to working directory /app
WORKDIR /app

# Copy the code into the container
COPY . .

RUN go mod download

RUN make build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin/app /bin/app

# Export necessary port
EXPOSE 80
# Command to run when starting the container
CMD ["/bin/app"]