FROM golang:tip-alpine3.22
WORKDIR /app
ENV GOPROXY=direct
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
EXPOSE 8080
CMD [ "go","run","."]
