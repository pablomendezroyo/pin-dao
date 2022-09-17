#!/bin/bash -x


## INPUTS(Github Repo, commit, ENS, IPFSNode, isChecker, PrivateKey)


## Operation functions
propose () {

# 1. Download github repo content of a determined commit
[ ! -d "./repo" ] && git clone ${REPO} repo

# # check commit exists()instead  clone master
cd repo
git checkout ${COMMIT}

# 2. Upload the content to IPFS
## Upload to IPFS and save hashes
ipfs --api=/dns/ipfs.dappnode/tcp/5001 add -p -r . --quiet | tee ../listHashes

#cat ../listHashes # for testing
IPFS_HASH_NO_BUILT=$(tail -1 ../listHashes)

# 3. Generate manifest

cat << EOF > "manifest.json"
{"GH_REPO": "$REPO", "COMMIT": "$COMMIT", "IPFS_HASH_REPO": "$IPFS_HASH_NO_BUILT", "ENS": "$ENS"}
EOF
cat ./manifest.json


# 4. Build the content

npm i -g yarn

yarn build:static

# 5. Update to IPFS

ipfs --api=/dns/ipfs.dappnode/tcp/5001 add -p -r . --quiet | tee ../listHashesBuilt

echo "IPFS hash built" 
#tail -1 ../listHashesBuilt # for testing
IPFS_HASH_BUILT=$(tail -1 ../listHashesBuilt)

curl    --connect-timeout 5 \
        --max-time 10 \
        --retry 5 \
        --retry-delay 0 \
        --retry-max-time 40 \
        -X POST "http://my.dappnode/data-send?key=IPFS_URL&data=http://${IPFS_HASH_BUILT}.ipfs.ipfs.dappnode:8080/" \
        || { echo "[ERROR] failed to post the UI_IPFS_HASH IPFS to dappmanager"; }

echo $(ls -a)
"Falta como acceder a una web en IPFS y no solo a sus archivos"

}

## Check what kind of operation will be executed:
## - checker 
## - proposer




if [ ${OPERATION} == "checker" ]; then
   echo "Check if the IPFS content mantains the integrity"
elif [ ${OPERATION} == 'proposer' ]
then
    echo "Start the process of proposing an repository content to the DAO"
    propose
else
    echo "Invalid operation: the value ${OPERATION} must be checker or proposer"
    exit 1
fi


# OPTION EXTRA A
# Checker(commit,repo, IPFS_HASH_BUILT)
## 1. Download files from the Hash IPFS (no-built)
## 2. Download files from the Hash IPFS (built)
## 3. Generating the manifest
## 4. Make the build of the not-built content
## 5. IPFS build
## 6. Check

# 1. Download github repo content of a determined commit
# [ ! -d "./repo" ] && git clone ${REPO} repo

# # check commit exists()instead  clone master
# cd repo
# git checkout ${COMMIT}

# Checker(IPFS_HASH_REPO, IPFS_HASH_BUILT)
## 1. Download files from the Hash IPFS (no-built)
## 2. Download files from the Hash IPFS (built)
## 3. Generating the manifest
## 4. Make the build of the not-built content
## 5. IPFS build
## 6. Check

## 1. Download the repo content from IPFS

#ipfs ls QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG

## 2. Download files from the Hash IPFS (built)

## 3. generating manifest (do it in jq or jo)

#echo "{ GH_REPO:$REPO,COMMIT:$COMMIT,IPFS_HASH_REPO:$VALUE3,ENS:$ENS }" > manifest.json


# Proposer Function
## 1. Download files from the github (no-built)
## 2. Update repo content to IPFS
## 3. Generating the manifest
## 4. Make the build of the not-built content
## 5. IPFS build
## 6. SC addENS


