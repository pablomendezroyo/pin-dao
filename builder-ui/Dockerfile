FROM node:18-bullseye-slim

WORKDIR /usr/src/app
RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

COPY --from=ipfs/kubo:v0.14.0 /usr/local/bin/ipfs /usr/local/bin/ipfs
COPY entrypoint.sh /usr/local/bin

CMD ["sleep 10000"]
#ENTRYPOINT [ "entrypoint.sh" ]