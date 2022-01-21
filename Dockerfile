FROM golang:1.17.2-alpine


WORKDIR /go/src/app/
COPY ./ .

RUN apk upgrade --update && \
    apk --no-cache add git gcc

RUN go get -u github.com/go-chi/chi/v5
RUN go mod tidy

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

RUN go get -u github.com/go-delve/delve/cmd/dlv && \
    go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

CMD ["air", "-c", ".air.toml"]