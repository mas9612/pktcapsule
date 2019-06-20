FROM golang:1.12.6-stretch AS build

LABEL maintainer="Masato Yamazaki <mas9612@gmail.com>"

RUN git clone https://github.com/mas9612/pktcapsule.git /pktcapsule
WORKDIR /pktcapsule
RUN make test && make build


FROM alpine:3.10.0

RUN mkdir /app
WORKDIR /app
COPY --from=build /pktcapsule/pktcapsuled .

EXPOSE 10000

ENTRYPOINT ["/app/pktcapsuled"]
