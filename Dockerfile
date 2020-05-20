FROM golang:1.7-onbuild

RUN mkdir /resume-api

ADD . /resume-api/
WORKDIR /resume-api

RUN go build .
EXPOSE 80
COPY ./datasource/*.csv /resume-api/
CMD ["/resume-api/resume-api"]
