FROM golang:1.13

RUN apt-get -y update
RUN apt-get -y install python3 python3-pip
RUN pip3 install Scrapy

COPY . /opt/

WORKDIR /opt
RUN go build -o osc ./server

WORKDIR /opt/crawler
RUN pip3 install -r requirements.txt

EXPOSE 8080

CMD ["/opt/osc"]