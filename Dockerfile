FROM golang:1.16-alpine

RUN mkdir /work 
ADD . /work
WORKDIR /work

RUN go mod download
RUN go build -o ./work/main .

EXPOSE 3000
CMD ["./work/main"]