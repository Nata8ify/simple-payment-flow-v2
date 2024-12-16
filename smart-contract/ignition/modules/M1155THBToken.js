// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("M1155THBTokenModule", (m) => {
  const m1155thbToken = m.contract("M1155THBToken", [], {});
  return { m1155thbToken };
});