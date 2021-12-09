FROM golang:latest

WORKDIR /application

COPY . . 

RUN go mod download

RUN go build -o app ./

ENTRYPOINT [ "./app" ]