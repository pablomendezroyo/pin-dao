FROM ubuntu

RUN apt update && apt install -y git wget curl

COPY --from=ipfs/kubo:v0.15.0 /usr/local/bin/ipfs /usr/local/bin/ipfs

COPY entrypoint.sh /usr/local/bin

CMD entrypoint.sh