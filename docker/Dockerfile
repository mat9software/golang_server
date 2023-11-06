# Usage:
# docker build --build-arg GO_VERSION=<version> -t ubuntu_go .
# docker build --build-arg GO_VERSION=1.21.3 -t ubuntu_go .


FROM ubuntu:latest

ARG GO_VERSION
ENV GO_VERSION=${GO_VERSION}

RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR /srv

RUN apt install -y cmake tmux

RUN apt install -y clang-format locate universal-ctags

RUN git config --global --add safe.directory /srv

RUN echo 'export EDITOR=vim'

RUN apt install -y gdb

RUN apt install -y nginx
RUN wget https://github.com/nginx/nginx/raw/master/conf/mime.types
RUN cp mime.types /etc/nginx/mime.types

RUN apt install -y vim curl

COPY vimrc /root/.vimrc
COPY vimrc /root/.gdbinit

RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | /bin/bash
ENV NVM_DIR=/root/.nvm
ENV NODE_VERSION=v16.14.0
RUN . $HOME/.nvm/nvm.sh && nvm install $NODE_VERSION

#RUN git config --global user.name ""
#RUN git config --global user.email ""

COPY nginx.conf /etc/nginx/sites-available/default

ENTRYPOINT service nginx restart && bash
