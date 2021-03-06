# we use a multi stage build to have a separate build and run env
FROM golang:latest as buildenv
LABEL maintainer="nicola.murino@gmail.com"
RUN go get -d github.com/drakkan/sftpgo
WORKDIR /go/src/github.com/drakkan/sftpgo
# uncomment the next line to get the latest stable version instead of the latest git
#RUN git checkout `git rev-list --tags --max-count=1`
RUN go build -i -ldflags "-s -w -X github.com/drakkan/sftpgo/utils.commit=`git describe --always --dirty` -X github.com/drakkan/sftpgo/utils.date=`date -u +%FT%TZ`" -o sftpgo

# now define the run environment
FROM debian:latest

ARG BASE_DIR=/app
ARG DATA_REL_DIR=data
ARG CONFIG_REL_DIR=config
ARG USERNAME=sftpgo
ARG GROUPNAME=sftpgo
ARG UID=515
ARG GID=515

# HOME_DIR for sftpgo itself
ENV HOME_DIR=${BASE_DIR}/${USERNAME}
# DATA_DIR, this is a volume that you can use hold user's home dirs
ENV DATA_DIR=${BASE_DIR}/${DATA_REL_DIR}
# CONFIG_DIR, this is a volume to persist the daemon private keys, configuration file ecc..
ENV CONFIG_DIR=${BASE_DIR}/${CONFIG_REL_DIR}

RUN mkdir -p ${DATA_DIR} ${CONFIG_DIR}
RUN groupadd --system -g ${GID} ${GROUPNAME}
RUN useradd --system --create-home --no-log-init --home-dir ${HOME_DIR} --comment "SFTPGo user" --shell /bin/false --gid ${GID} --uid ${UID} ${USERNAME}

WORKDIR ${HOME_DIR}
RUN mkdir -p bin .config/sftpgo
ENV PATH ${HOME_DIR}/bin:$PATH
COPY --from=buildenv /go/src/github.com/drakkan/sftpgo/sftpgo bin/sftpgo
# default config file to use if no config file is found inside the CONFIG_DIR volume.
# You can override each configuration options via env vars too
COPY --from=buildenv /go/src/github.com/drakkan/sftpgo/sftpgo.json .config/sftpgo/
RUN chown -R ${UID}:${GID} ${DATA_DIR}

# run as non root user
USER ${USERNAME}

EXPOSE 2022 8080

# the defined volumes must have write access for the UID and GID defined above
VOLUME [ "$DATA_DIR", "$CONFIG_DIR" ]

# override some default configuration options using env vars
ENV SFTPGO_CONFIG_DIR=${CONFIG_DIR}
# setting SFTPGO_LOG_FILE_PATH to an empty string will log to stdout
ENV SFTPGO_LOG_FILE_PATH=${CONFIG_DIR}/sftpgo.log
ENV SFTPGO_HTTPD__BIND_ADDRESS=""

ENTRYPOINT ["sftpgo"]
CMD ["serve"]