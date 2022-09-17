############################
# Get the abi from the sc ##
############################
FROM ethereum/solc:0.8.17-alpine as abi
WORKDIR /usr/src/app
COPY ./sc/contracts/PinnerDao.sol .
RUN solc --abi PinnerDao.sol -o build

##################################
## Generate the sc in go format ##
##################################
FROM ethereum/client-go:alltools-latest as go-sc
WORKDIR /usr/src/app
COPY --from=abi /usr/src/app/build/TempNameDao.abi ./
RUN abigen --abi=/usr/src/app/TempNameDao.abi --pkg=main --out=PinnerDao.go

###############################
## Build the go docker image ##
###############################
FROM golang:1.17-alpine 
WORKDIR /usr/src/app

# Needed to create binary to be executed in debian. See https://pkg.go.dev/cmd/cgo
ENV CGO_ENABLED=0
COPY ./ipfs/pinner-dao/src/ ./ 
COPY --from=go-sc /usr/src/app/PinnerDao.go /usr/src/app/
COPY --from=go-sc /usr/src/app/TempNameDao.abi /usr/src/app/
RUN go build
ENTRYPOINT ["/usr/src/app/pinner-dao"]
