FROM golang:1.13

RUN apt-get -y update
RUN apt-get -y install python3 python3-pip
RUN pip3 install Scrapy

COPY ./crawler /opt/crawler
COPY ./server /opt/server

WORKDIR /opt
RUN go build -o osc-server ./server

WORKDIR /opt/crawler
RUN pip3 install -r requirements.txt

RUN mkdir -p /opt/data

EXPOSE 8080

CMD ["/opt/osc-server"]