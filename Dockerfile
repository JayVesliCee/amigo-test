FROM golang:1.8

MAINTAINER Jean-Charles PASSEPONT "jc.epitech@gmail.com"

WORKDIR /go/src/app
COPY . .


EXPOSE 8080

RUN go build -o amigo
CMD ["./amigo"]
