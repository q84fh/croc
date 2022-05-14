FROM golang:alpine

RUN mkdir /code

ADD croc.go /code
ADD croc_awake.html /code
ADD croc_sleeping.html /code

WORKDIR /code

RUN go mod init github.com/q84fh/croc/v2
RUN go mod tidy
RUN go build -o croc .

CMD ["/code/croc"]
