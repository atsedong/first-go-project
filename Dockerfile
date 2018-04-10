FROM golang:latest as builder

WORKDIR /go/src/github.com/first-go-project

RUN go get -d -v golang.org/x/net/html
#RUN go get -d -v ./...

COPY main.go .
COPY index.html .
COPY /static ./static/

# Compile our application to produce a static binary called app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Copy our compiled binary to an alpine image
FROM alpine:latest

COPY --from=builder /go/src/github.com/first-go-project ./

CMD ["./app"]
