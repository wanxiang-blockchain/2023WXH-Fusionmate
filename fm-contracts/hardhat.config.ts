import 'dotenv/config';

import 'chai';
import { getMnemonic, getNodeUrl } from './utils/network';
import { HardhatUserConfig } from 'hardhat/config';
import 'hardhat-gas-reporter'; // can get tx gas info
import '@nomiclabs/hardhat-etherscan'; // verify source code
//https://github.com/wighawag/hardhat-deploy#npm-install-hardhat-deploy
import '@nomiclabs/hardhat-ethers';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';

const config: HardhatUserConfig = {
    defaultNetwork: 'hardhat',
    networks: {
        hardhat: {
            deploy: ['deploy_mock', 'deploy_testnet'],
            saveDeployments: true,
            allowUnlimitedContractSize: true,
        },
        goerli: {
            deploy: ['deploy_testnet'],
            url: getNodeUrl('goerli'),
            accounts: {
                mnemonic: getMnemonic('goerli'),
                count: 20,
                // accounts: process.env.PRIVATE_KEY !== undefined ? [process.env.PRIVATE_KEY] : [],
            },
            gas: 'auto',
        },
    },
    gasReporter: {
        enabled: process.env.REPORT_GAS !== undefined,
        // enabled: true,
        currency: 'USD',
    },
    etherscan: {
        apiKey: process.env.ETHERSCAN_API_KEY,
    },
    namedAccounts: {
        deployer: 0,
        trader: 1,
    },
    solidity: {
        compilers: [
            {
                version: '0.8.13',
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 100,
                    },
                },
            },
        ],
    },
};

export default config;
