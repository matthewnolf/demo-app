FROM golang:latest

WORKDIR /application

COPY . . 

RUN go mod download

RUN go build -ldflags "-X main.gitRevision=$(git rev-list -1 HEAD)" -o app ./

ENTRYPOINT [ "./app" ]