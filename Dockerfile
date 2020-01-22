FROM python:3-alpine as python

RUN pip3 install Jinja2 pyyaml
WORKDIR /app

COPY ./templates/ ./templates/
COPY generate.py ./
COPY config.yaml ./
RUN ls ./

RUN python3 generate.py

FROM golang:alpine as go

workdir /app

ARG SOAP_ENDPOINT=http://www.dneonline.com/calculator.asmx\?WSDL

RUN apk add --no-cache git
RUN go get github.com/hooklift/gowsdl/...

RUN gowsdl -o gen -p gen ${SOAP_ENDPOINT}

COPY ./ ./

RUN go build

RUN ls

FROM alpine

COPY --from=go /app/soap2rest /app/soap2rest


ENTRYPOINT /app/soap2rest


