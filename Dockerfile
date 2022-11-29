## Build Image
FROM golang:1.18.1-alpine as build

# this comes from standard alpine nightly file
#  https://github.com/rust-lang/docker-rust-nightly/blob/master/alpine3.12/Dockerfile
# with some changes to support our toolchain, etc
RUN set -eux; apk add --no-cache ca-certificates build-base;
RUN apk add git

WORKDIR /nova
COPY . /nova

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 9ecb037336bd56076573dc18c26631a9d2099a7f2b40dc04b6cae31ffb4c8f9a
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 6e4de7ba9bad4ae9679c7f9ecf7e283dd0160e71567c6a7be6ae47c81ebe7f32
RUN cp /lib/libwasmvm_muslc.$(uname -m).a /lib/libwasmvm_muslc.a
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build

## Deploy image
FROM golang:1.18.1-alpine

COPY --from=build /nova/build/novad /bin/novad

ENV HOME /nova
WORKDIR $HOME

EXPOSE 26656 
EXPOSE 26657
EXPOSE 1317

USER nova
