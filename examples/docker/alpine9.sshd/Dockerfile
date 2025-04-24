FROM alpine:3.20.2
ENTRYPOINT ["/entrypoint.sh"]
COPY rootfs /

RUN wget -qO- https://gitee.com/clh21/sh/raw/master/mirror.sh | sh

RUN apk add --no-cache openssh \
  && sed -i s/#PermitRootLogin.*/PermitRootLogin\ yes/ /etc/ssh/sshd_config \
  && echo "root:root" | chpasswd

EXPOSE 22
