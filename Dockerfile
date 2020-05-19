FROM golang:1.6-onbuild

RUN mkdir /resume-api

ADD . /resume-api/
WORKDIR /resume-api

RUN go build .

CMD ["/resume-api/resume"]