# sshutil

## prepare test host
```bash
docker run -d --rm -p 8022:22 leehom/detect:alpine9.sshd
# images build by dockerfile with /examples/docker/alpine9.sshd/Dockerfile use build.sh
# test host ssh login ok
ssh root@localhost -p 8022
# password is root
```

## do test

```bash
make run
```