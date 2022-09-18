# IPFS pinner

The IPFS pinner is automatically grabbing the SC of the DAO from `./sc/contracts/PinnerDao.sol`, compile it and creates the corresponding ABI and the smart contract in go code required by the pinner.


### RUN

To build the IPFS pinner run the following command at the root project directory:
```
docker-compose -f docker-compose-pinner.yml build
```
To run the IPFS pinner run the following command at the root project directory:
```
docker-compose -f docker-compose-pinner.yml up
```

### ENV
If the `ENVIRONMENT=development` use the `dev.env`. If the `ENVIRONMENT=production` use the `prod.env`. The environment required are an IPFS API endpoint and an Ethereum RPC from Rinkeby testnet.

```
IPFS_API_URL=
ETH_RPC_URL=
```
