FROM alpine
RUN apk add netcat-openbsd curl
RUN apk add docker-cli

EXPOSE 19030

ARG WORKDIR=/usr/yunji/cloudiac/
WORKDIR ${WORKDIR}

COPY configs/config-runner.yml.sample config-runner.yml
COPY build/ct-runner ct-runner

## runner 容器化部署时不能设置 assets_path
RUN sed -ie 's/ assets_path:.*/ assets_path: ""/' config-runner.yml

CMD ./ct-runner
