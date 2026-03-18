# smarter-screen
Tech screen for Smarter Technologies

## Prerequisites

Install Go (1.23+): https://go.dev/doc/install

Verify the installation:

```sh
go version
```

## Setup

Clone the repository and download dependencies:

```sh
git clone https://github.com/elkirk/smarter-screen
cd smarter-screen
go mod tidy
```

## Running

The `Sort` function takes four arguments: `width`, `height`, `length` (in cm) and `mass` (in kg).

### With `go run`

```sh
go run cmd/main.go <width> <height> <length> <mass>
```

Example:

```sh
go run cmd/main.go 50 50 50 10
# STANDARD
```

### Compile to binary

```sh
go build -o sort cmd/main.go
./sort <width> <height> <length> <mass>
```

Example:

```sh
go build -o sort cmd/main.go
./sort 150 100 100 25
# REJECTED
```

## Running Tests

```sh
go test -v ./...
```
