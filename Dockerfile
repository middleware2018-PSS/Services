FROM golang:1.9.4 as builder
WORKDIR /go/src/app
RUN go get -u github.com/golang/dep/cmd/dep
ADD Gopkg.toml Gopkg.lock /go/src/app/
ADD src /go/src/app/
RUN dep ensure 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app/app ./
ADD ./src/config config
EXPOSE 5000
CMD ["./app"]  

