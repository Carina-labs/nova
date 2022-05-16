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
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-rc.0/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep eda70adcd2f09a0ae0a6c1e53ecc318809e6d57244bda7b7a29ccd9cf591aa37  

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.a

RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build

## Deploy image
FROM golang:1.18.1-alpine

COPY --from=build /nova/build/novad /bin/novad

ENV HOME /nova
WORKDIR $HOME

EXPOSE 26656 
EXPOSE 26657
EXPOSE 1317
