FROM daocloud.io/library/golang:latest

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . /usr/src/app
RUN echo 'data'
RUN ls
RUN echo 'now directory'
RUN pwd
RUN export GOPATH=/usr/src/app
RUN echo 'go path'
RUN $GOPATH
RUN go get github.com/go-martini/martini
RUN go get github.com/martini-contrib/render
RUN cd src
RUN go build
RUN mv hello ../
RUN cd ..
RUN ls

EXPOSE 3000
CMD ./hello