FROM golang:latest


WORKDIR /go/src/app/
COPY ./ .

RUN go get -u github.com/go-chi/chi/v5
RUN go mod tidy

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@latest