require("@nomicfoundation/hardhat-toolbox");
var vars = require("hardhat/config").vars;

/** @type string */

const CONTRACT_MTHB_NODE_ENDPOINT = vars.get("CONTRACT_MTHB_NODE_ENDPOINT")
const CONTRACT_MTHB_OWNER_CONTRACT_PK = (vars.has("CONTRACT_MTHB_CHAIN_ID") ? parseInt(vars.get("CONTRACT_MTHB_CHAIN_ID")) : 1337)
const CONTRACT_MTHB_CHAIN_ID = vars.get("CONTRACT_MTHB_OWNER_CONTRACT_PK")

module.exports = {
    solidity: "0.8.28",
    networks: {
        besu: {
            url: CONTRACT_MTHB_NODE_ENDPOINT,
            chainId: CONTRACT_MTHB_OWNER_CONTRACT_PK,
            accounts: [CONTRACT_MTHB_CHAIN_ID]
        }
    }
};
