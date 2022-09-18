# Pin DAO

## IPFS

DAO members will be able to propose new websites source code content and build files to be pinned. These websites will be served in a node of IPFS, and can be consume through any IPFS gateway.

Creating an incentive program for the pinners to pin the content defined by the DAO is out of the scope of this project due to its difficulty to truly ensure that a given pinners is really pinning any given content. This is another huge project that Filecoin is working on.

Pinners are subscribed to the SC of the DAO that returns the ENS under where there are the websites to be pinned. The pinner resolves each ENS to the corresponding IPFS hash and pins its content. Furthermore, if the content hash contains the file `manifest.json` it will try to get the original Github repository so it get pinned for resilience.

## BUILDER-UI

The Builder-UI is a toolkit that have mainly two options: **proposer**
## DAO

## Monitoring UI


### Future

In order to move forward the project and simplify it, the builder-ui, ipfs-pinner and monitoring-ui will be unified into a single user interface that:
- Will monitor the SC of the DAO, letting users know which ENS are getting pinned.
- Will allow to run a pinner in the background to contribute the DAO pinning system
- Will allow to propose a hash to a given ENS, by running the builder-ui toolkit in the background. Furthermore it will be allowed to validate a proposal of the DAO for a given hash with the builder-ui toolkit, so it gets compared with the hash output and DAO users can ensure if is valid or not.
