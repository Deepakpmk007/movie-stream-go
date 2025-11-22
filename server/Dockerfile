FROM golang:tip-alpine3.22

RUN apk add --no-cache git curl

# Install Air once (cached)
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD [ "air" ]