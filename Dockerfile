## Build Image
FROM golang:1.18.1-alpine as build

WORKDIR /novachain
COPY . /novachain

# From https://github.com/CosmWasm/wasmd/blob/master/Dockerfile
# For more details see https://github.com/CosmWasm/wasmvm#builds-of-libwasmvm 
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-beta7/libwasmvm_muslc.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep d0152067a5609bfdfb3f0d5d6c0f2760f79d5f2cd7fd8513cafa9932d22eb350
RUN apk update && apk add --no-cache gcc libc-dev git make
RUN BUILD_TAGS=muslc make build

## Deploy image
FROM golang:1.18.1-alpine

COPY --from=build /novachain/build/novachaind /bin/novachaind

ENV HOME /novachain
WORKDIR $HOME

EXPOSE 26656 
EXPOSE 26657
EXPOSE 1317
