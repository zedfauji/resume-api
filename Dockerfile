FROM golang:1.7-onbuild

RUN mkdir /resume-api

ADD . /resume-api/
WORKDIR /resume-api

RUN go build .

CMD ["/resume-api/resume-api"]