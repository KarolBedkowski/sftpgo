FROM golang:1.13-alpine3.10 as builder

RUN apk add --no-cache git gcc g++ ca-certificates \
  && go get -d github.com/drakkan/sftpgo
WORKDIR /go/src/github.com/drakkan/sftpgo
# uncomment the next line to get the latest stable version instead of the latest git
#RUN git checkout `git rev-list --tags --max-count=1`
RUN go build -i -ldflags "-s -w -X github.com/drakkan/sftpgo/utils.commit=`git describe --always --dirty` -X github.com/drakkan/sftpgo/utils.date=`date -u +%FT%TZ`" -o /go/bin/sftpgo

FROM alpine:3.10

RUN  apk add --no-cache ca-certificates su-exec \
  && mkdir -p /data /etc/sftpgo /srv/sftpgo/config

COPY --from=builder /go/bin/sftpgo /bin/
COPY --from=builder /go/src/github.com/drakkan/sftpgo/sftpgo.json /etc/sftpgo/sftpgo.json
COPY docker-entrypoint.sh /bin/entrypoint.sh
RUN chmod +x /bin/entrypoint.sh

VOLUME [ "/data", "/srv/sftpgo/config" ]
EXPOSE 2022 8080

ENTRYPOINT ["/bin/entrypoint.sh"]
CMD []