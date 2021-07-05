FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
#RUN go get gitlab.com/claude.roy790/main
#RUN cd /build && git clone https://gitlab.com/claude.roy790/docker-test.git

COPY . /build/

RUN cd /build/main && go build

EXPOSE 4000

CMD [ "./main/docker-test" ]