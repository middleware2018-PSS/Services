FROM golang:1.9.4 as builder
WORKDIR /go/src/github.com/middleware2018-PSS/Services
RUN go get -u github.com/golang/dep/cmd/dep
ADD Gopkg.toml Gopkg.lock /go/src/github.com/middleware2018-PSS/Services/
RUN dep ensure --vendor-only
ADD . /go/src/github.com/middleware2018-PSS/Services/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/middleware2018-PSS/Services/build/app ./
ADD config config
EXPOSE 5000
CMD ["./app"]  

