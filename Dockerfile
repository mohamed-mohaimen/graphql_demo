FROM golang:1.15-alpine

RUN apk add --no-cache git
RUN apk add --update make
RUN mkdir -p /app

# Move to working directory /app
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

RUN go build -o bin/app

# Export necessary port
EXPOSE 9090
# Command to run when starting the container
CMD ["bin/app"]