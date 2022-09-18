#!/bin/bash

set -x

cd /usr/src/app

exec REACT_APP_DAPPNODE_URL=https://ipfs.${_DAPPNODE_GLOBAL_DOMAIN} npm start