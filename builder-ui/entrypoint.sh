#!/bin/bash -x

## Operation functions

# Propose Function
## 1. Download files from the github (no-build)
## 2. Update repo content to IPFS
## 3. Generating the manifest
## 4. Make the build of the not-build content
## 5. IPFS build TODO its required to show the UI not only se it 
source ~/.nvm/nvm.sh

propose () {

# 1. Download github repo content of a determined commit
[ ! -d "./repo" ] && git clone ${REPO} repo

# # check commit exists()instead  clone master
cd repo
git checkout ${COMMIT}

# 2. Upload the content to IPFS
## Upload to IPFS and save hashes
ipfs --api=/dns/ipfs.dappnode/tcp/5001 add --hidden -p -r . --quiet | tee ../listHashes

#cat ../listHashes # for testing
IPFS_HASH_NO_BUILD=$(tail -1 ../listHashes)

# 3. Generate manifest

cat << EOF > "manifest.json"
{"GH_REPO": "$REPO", "COMMIT": "$COMMIT", "IPFS_HASH_REPO": "$IPFS_HASH_NO_BUILD", "ENS": "$ENS"}
EOF
cat ./manifest.json
CHECKSUM_REPO=$(tar cvf - /repo | sha1sum)
# 4. Build the content

npm i -g yarn

yarn

yarn build:static

# 5. Update to IPFS
echo "directorio tras el build"
echo $(ls -a)
CHECKSUM_BUILD_NO_MANIFEST=$(tar cvf - ./out | sha1sum )


cp ./manifest.json ./out/manifest.json
ipfs --api=/dns/ipfs.dappnode/tcp/5001 add -p -r ./out --quiet | tee ../listHashesbuild

#echo "IPFS hash build" 
#cat ../listHashesbuild
 # for testing
echo "IPFS_HASH_BUILD:"
tail -1 ../listHashesbuild
IPFS_HASH_BUILD=$(tail -1 ../listHashesbuild)
echo "IPFS_HASH_CODED:"
IPFS_HASH_CODED=$(ipfs cid base32 $IPFS_HASH_BUILD)
echo "http://${IPFS_HASH_CODED}.ipfs.ipfs.dappnode:8080/"


CHECKSUM_BUILD=$(tar cvf - ./out | sha1sum )
echo "Checksum repo $CHECKSUM_REPO"
echo "Checksum build $CHECKSUM_BUILD"
echo "Checksum build no manifest $CHECKSUM_BUILD_NO_MANIFEST"
}

# Checker(IPFS_HASH_REPO, IPFS_HASH_build)
## 1. Download files from the Hash IPFS (no-build)
## 2. Download files from the Hash IPFS (build)
## 3. Generating the manifest
## 4. Make the build of the not-build content
## 5. IPFS build
## 6. Check

validate () {
    echo "  Validation process"
    ipfs --api=/dns/ipfs.dappnode/tcp/5001 get $IPFS_HASH_BUILD  --output=/build_repo_content
    CHECKSUM_REPO=$(tar cvf - /initial_repo_content/out | sha1sum)
    echo "ls of the build repo content $(ls -a /build_repo_content)"
    ipfs --api=/dns/ipfs.dappnode/tcp/5001 get $IPFS_HASH_REPO  --output=/initial_repo_content
    echo $(ls -a /initial_repo_content)

    ## 3. Generate manifest
    echo "Generate manifest"
    cat << EOF > "manifest.json"
    {"GH_REPO": "$REPO", "COMMIT": "$COMMIT", "IPFS_HASH_REPO": "$IPFS_HASH_NO_BUILD", "ENS": "$ENS"}
EOF

cat ./manifest.json
    ## 4. make the build
    cd ./initial_repo_content

    npm i -g yarn
    yarn
    yarn build:static

    ## add manifest
    mv ../manifest.json ./out/manifest.json

    cd ..

    ## 6. Calculate checksums + Check values
    
    CHECKSUM_BUILD=$(tar cvf - /build_repo_content | sha1sum )
    echo "Checksum repo $CHECKSUM_REPO"
    echo "Checksum build $CHECKSUM_BUILD"
}

## Check what kind of operation will be executed:
## - validate 
## - proposer

if [ ${OPERATION} == "checker" ]; then
    echo "Validate if the IPFS content mantains the integrity"
    validate
elif [ ${OPERATION} == 'propose' ]
then
    echo "Start the process of proposing an repository content to the DAO"
    propose
else
    echo "Invalid operation: the value ${OPERATION} must be checker or proposer"
    exit 1
fi


# OPTION EXTRA A
# Checker(commit,repo, IPFS_HASH_build)
## 1. Download files from the Hash IPFS (no-build)
## 2. Download files from the Hash IPFS (build)
## 3. Generating the manifest
## 4. Make the build of the not-build content
## 5. IPFS build
## 6. Check

# 1. Download github repo content of a determined commit
# [ ! -d "./repo" ] && git clone ${REPO} repo

# # check commit exists()instead  clone master
# cd repo
# git checkout ${COMMIT}

## 1. Download the repo content from IPFS

#ipfs ls QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG

## 2. Download files from the Hash IPFS (build)

## 3. generating manifest (do it in jq or jo)

#echo "{ GH_REPO:$REPO,COMMIT:$COMMIT,IPFS_HASH_REPO:$VALUE3,ENS:$ENS }" > manifest.json





