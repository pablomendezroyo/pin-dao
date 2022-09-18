#!/bin/bash

WORKING_DIR=./build
rm -rf $WORKING_DIR
mkdir -p $WORKING_DIR

propose() {
    # Read input from the user to set the above inputs
    read -p "Enter the github repo url: " REPO
    read -p "Enter the build output dir of the given repository: " BUILD_DIR
    read -p "Enter the commit hash: " COMMIT
    read -p "Enter the ENS name: " ENS
    read -p "Enter the IPFS api url (i.e ipfs.dappnode): " IPFS_URL

    # Download github repo content of a determined commit
    git clone "${REPO}" ${WORKING_DIR}
    cd ${WORKING_DIR}
    git checkout "${COMMIT}"

    # Upload the content to IPFS and save hashes
    ipfs --api=/dns/"${IPFS_URL}"/tcp/5001 add --hidden -p -r . --quiet | tee ../listHashes

    # cat ../listHashes # for testing
    IPFS_HASH_REPO=$(tail -1 ../listHashes)

    yarn
    yarn build:static

    # Generate manifest_dappnode
    cat <<EOF >"manifest_dappnode.json"
{"GH_REPO": "$REPO", "COMMIT": "$COMMIT", "IPFS_HASH_REPO": "$IPFS_HASH_REPO", "ENS": "$ENS"}
EOF
    mv manifest_dappnode.json "./${BUILD_DIR}"

    # Upload the build directory
    ipfs --api=/dns/"${IPFS_URL}"/tcp/5001 add -p -r ./"${BUILD_DIR}" --quiet | tee ../listHashesbuild

    echo "IPFS_HASH_BUILD:"
    tail -1 ../listHashesbuild
    IPFS_HASH_BUILD=$(tail -1 ../listHashesbuild)
    echo "IPFS_HASH_CODED:"
    IPFS_HASH_CODED=$(ipfs cid base32 $IPFS_HASH_BUILD)
    echo "http://${IPFS_HASH_CODED}.ipfs.ipfs.dappnode:8080/"

}

validate() {
    read -p "Enter the build output dir of the given repository: " BUILD_DIR
    read -p "Enter the IPFS api url (i.e ipfs.dappnode): " IPFS_URL
    read -p "Enter the IPFS hash of the build directory: " IPFS_HASH_BUILD
    read -p "Enter the IPFS hash of the repo directory: " IPFS_HASH_REPO

    cd ${WORKING_DIR}

    ipfs --api=/dns/ipfs.dappnode/tcp/5001 get $IPFS_HASH_BUILD --output=./build_repo_content
    echo "ls of the build repo content $(ls -a /build_repo_content)"
    ipfs --api=/dns/ipfs.dappnode/tcp/5001 get $IPFS_HASH_REPO --output=./initial_repo_content
    echo $(ls -a /initial_repo_content)

    cd ./initial_repo_content

    npm i -g yarn
    yarn
    yarn build:static

    # Generate manifest_dappnode
    echo "Generate manifest_dappnode"
    cat <<EOF >"manifest_dappnode.json"
    {"GH_REPO": "$REPO", "COMMIT": "$COMMIT", "IPFS_HASH_REPO": "$IPFS_HASH_REPO", "ENS": "$ENS"}
EOF
    mv manifest_dappnode.json initial_repo_content/"${OUT_DIR}"

    ## 6. Calculate checksums + Check values

    #CHECKSUM_REPO=$(shasum -a256 "/initial_repo_content/out")
    #CHECKSUM_BUILD=$(shasum -a256 "/build_repo_content")
    #echo "Checksum repo $CHECKSUM_REPO"
    #echo "Checksum build $CHECKSUM_BUILD"
}

read -p "Enter the operation to be performed (validate or propose): " OPERATION

if [ ${OPERATION} == "checker" ]; then
    echo "Validate if the IPFS content mantains the integrity"
    validate
elif [ ${OPERATION} == 'propose' ]; then
    echo "Propose a decentralliced web3site to be added to the DAO and pinned"
    propose
else
    echo "Invalid operation: the value ${OPERATION} must be checker or proposer"
    exit 1
fi
