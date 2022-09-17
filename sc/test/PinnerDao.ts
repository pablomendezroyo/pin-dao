import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("TempNameDao", function () {
  describe("Read", function () {
    it("Should returns the list of ens domains", async function () {
        const [owner] = await ethers.getSigners();
        const TempNameDao = await ethers.getContractFactory("TempNameDao");
        const tempNameDao = await TempNameDao.deploy();
        const ensList = await tempNameDao.getENSList();
        expect(ensList.length).to.equal(1);
        expect(ensList[0]).to.equal("subdomain1.hehehehehe.eth");
    });
  });
});
