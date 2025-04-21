FROM golang:1.24 as builder

ENV LANG en_US.utf8
WORKDIR /dac

COPY . .
RUN make

FROM ubuntu:22.04 as base

WORKDIR /dac
COPY --from=builder /dac/bin/dac /usr/local/bin/
ENTRYPOINT ["dac"]
