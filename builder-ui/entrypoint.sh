#!/bin/bash -x


## INPUTS(Github Repo, commit, ENS, IPFSNode, isChecker, PrivateKey)

## Check what kind of operation will be executed:
## - checker 
## - proposer

if [ ${OPERATION} == "checker" ]; then
   echo "Check if the IPFS content mantains the integrity"
elif [ ${OPERATION} == 'proposer' ]
then
    echo "Start the process of proposing an repository content to the DAO"
else
    echo "Invalid operation: the value ${OPERATION} must be checker or proposer"
    exit 1
fi



# Checker(IPFS_HASH_NO_BUILT, IPFS_HASH_BUILT)
## 1. Download files from the Hash IPFS (no-built)
## 2. Download files from the Hash IPFS (built)
## 3. Generating the manifest
## 4. Make the build of the not-built content
## 5. IPFS build
## 6. Check


[ ! -d "./repo" ] && git clone --bare ${REPO} repo

cd repo
git checkout ${COMMIT}





# Proposer Function
## 1. Download files from the Hash IPFS (no-built)
## 2. Download files from the Hash IPFS (built)
## 3. Generating the manifest
## 4. Make the build of the not-built content
## 5. IPFS build
## 6. SC addENS

