FROM golang:1.21.0-alpine3.18
RUN apk update
RUN mkdir /maptalk
WORKDIR /maptalk
COPY . /maptalk/
RUN go mod download
CMD [ "go", "run", "/maptalk/cmd/maptalk/main.go" ]