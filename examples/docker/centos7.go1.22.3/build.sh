#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
cd "$( dirname "${BASH_SOURCE[0]}" )" || exit
docker build . -t leehom/detect:centos7.go1.22.3
