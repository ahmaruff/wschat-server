FROM golang:latest

# RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

# RUN go mod tidy
ENV GO111MODULE=off

RUN go build -o binary

# ENTRYPOINT ["/app/binary"]

CMD ["/app/binary"]