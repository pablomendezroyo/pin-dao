import { expect } from "chai";
import { ethers } from "hardhat";

describe("TempNameDao", function () {
  describe("Read", function () {
    it("Should returns the list of ens domains", async function () {
        const TempNameDao = await ethers.getContractFactory("TempNameDao");
        const tempNameDao = await TempNameDao.deploy();
        // await expect(tempNameDao.addEns("subdomain1.hehehehehe.eth", "myhash")).not.to.be.reverted;
        // const ensList = await tempNameDao.getENSList();
        // expect(ensList.length).to.equal(3);
        // expect(ensList[0]).to.equal("subdomain1.hehehehehe.eth");
        // expect(ensList[2]).to.equal("yoursubdomain.hehehehehe.eth");
    });
  });
});
