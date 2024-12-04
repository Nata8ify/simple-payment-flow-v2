// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const INITIAL_SUPPLY = 6_000_000_000_000_000_000_000_000_000_000_000n;

module.exports = buildModule("MTHBTokenModule", (m) => {
  const mthbToken = m.contract("MTHBToken", [INITIAL_SUPPLY], {});
  return { mthbToken };
});
