FROM ubuntu

ARG GOVERSION=go1.14.2

RUN cd /tmp && \
	apt-get -qq update && \
	apt-get -y install wget git vim && \
	apt-get -y install gcc build-essential perl && \
	wget http://mirrors.ustc.edu.cn/golang/$GOVERSION.linux-amd64.tar.gz && \
	cd /opt && \
	tar zxf /tmp/$GOVERSION.linux-amd64.tar.gz && \
	rm -f /tmp/$GOVERSION.linux-amd64.tar.gz && \
	mkdir -p /opt/gopath/src && \
	mkdir -p /opt/gopath/pkg && \
	mkdir -p /opt/gopath/bin && \
	mkdir -p /opt/gopath/pkg/mod && \
	chgrp -R root /opt/gopath && \
	chmod g+rw /opt/gopath && \
	apt-get -y autoremove && \
	apt-get clean

ENV GOPATH=/opt/gopath
ENV GOROOT=/opt/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

