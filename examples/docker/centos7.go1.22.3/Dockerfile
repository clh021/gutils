#docker build . -t leehom/detect:centos7.go1.22.3
#docker run -it --rm -v `pwd`:/app -w /app -e CGO_ENABLED=0 -u ${UID}:${GID} leehom/detect:centos7.go1.22.3 bash
# FROM openanolis/anolisos:8.2-x86_64 # glibc 2.28
FROM centos:7
# glibc 2.17

MAINTAINER clh021@gmail.com

ENV GOLANG_VERSION go1.22.3

USER root

# no need with epel-release gcc
RUN yum install -y wget && \
    yum clean all && \
    rm -rf /var/cache/yum/*

# RUN wget -c https://golang.google.cn/dl/$GOLANG_VERSION.linux-amd64.tar.gz -O ~/go.tar.gz && \
RUN wget -c https://mirrors.aliyun.com/golang/$GOLANG_VERSION.linux-amd64.tar.gz -O ~/go.tar.gz && \
    tar -C /usr/local -xzf ~/go.tar.gz && \
    rm ~/go.tar.gz

ENV PATH /usr/local/go/bin:$PATH
ENV GOPROXY https://goproxy.cn,direct

RUN mkdir /.cache && \
    chmod 777 /.cache
