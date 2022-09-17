# Builder UI

The tool let you to run two types of operations:

- Checking: checking operation consists on check the integrity of the content of a IPFS hash, to use the checking operation define the value of env var called OPERATION in the docker-compose as checker. The inputs of checking operation are: Github Repo, commit, ENS, IPFSNode, PrivateKey
- Proposing: proposing operation consists on sent to the smart contact the ipfs hash of the content and a manifest with the required data, to use the proposing operation define the value of env var called OPERATION in the docker-compose as proposer. The inputs of proposing operation are: Github Repo, commit, ENS, IPFSNode, PrivateKey

`docker-compose build .`

`docker-compose run`
