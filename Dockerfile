FROM golang:1.13 as build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go vet ./...

RUN go get -u golang.org/x/lint/golint

RUN golint -set_exit_status ./...

RUN go build -o kafmesh-example ./cmd/kafmesh-example

FROM build as unit

CMD go test -race -coverprofile=/artifacts/coverage.txt -covermode=atomic ./...

FROM ubuntu:18.04 as final

WORKDIR /app

RUN apt update && apt install -y tzdata && apt clean

COPY --from=0 /build/kafmesh-example /app/

CMD ["/app/kafmesh-example"]

# testing
FROM build as endtoend

WORKDIR /build

RUN go get github.com/cucumber/godog/cmd/godog

WORKDIR /build/testing

CMD godog --strict --format=pretty /build/docs/features/
