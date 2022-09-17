import { ethers } from "hardhat";

async function main() {
  const TempNameDao = await ethers.getContractFactory("TempNameDao");
  const tempNameDao = await TempNameDao.deploy();

  await tempNameDao.deployed();

  console.log("TempNameDao deployed to:", tempNameDao.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
