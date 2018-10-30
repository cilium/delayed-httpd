FROM ubuntu
WORKDIR /
ADD . /
ENTRYPOINT ["/delayed-httpd"]
