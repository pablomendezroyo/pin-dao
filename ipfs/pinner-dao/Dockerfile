FROM ethereum/solc:0.8.17-alpine as abi

COPY ./sc/pinner-dao.sol .
RUN solc --abi Store.sol -o build

FROM ethereum/client-go:alltools-latest as go-sc
COPY from=abi /build/Store.abi /build/Store.abi/
RUN abigen --abi=./build/Store.abi --pkg=store --out=Store.go



