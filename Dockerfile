FROM daocloud.io/library/golang:latest

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY . /usr/src/app
RUN echo 'data'
RUN ls

EXPOSE 3000
CMD hello