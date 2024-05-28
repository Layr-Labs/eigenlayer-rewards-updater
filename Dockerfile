FROM golang:1.22-bullseye as build

RUN apt-get update
RUN apt-get install -y make

RUN mkdir /build

COPY . /build

WORKDIR /build

RUN make deps

RUN make

FROM golang:1.22-bullseye as run

COPY --from=build /build/bin/* /bin

ENTRYPOINT ["/bin/eigenlayer-rewards-updater"]
