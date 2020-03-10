FROM golang:1.14

# Check ulimit is good - this run line is here for debugging
RUN ulimit -l

WORKDIR /go/src/app

# Copy over the files for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy remaining source code
COPY . .

RUN go test -v

# Run tests
CMD ["go", "test", "-v"]
