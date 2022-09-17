// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

contract TempNameDao {
    constructor() {

    }

    function getENSList() public pure returns (string[] memory) {
        string[] memory ensList = new string[](1);
        ensList[0] = "subdomain1.hehehehehe.eth";
        return (ensList);
    }
}