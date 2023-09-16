import { BigNumber } from 'ethers';
import { AssistantType } from './types';

export const MAX_UINT256 = BigNumber.from('2')
  .pow(BigNumber.from('256'))
  .sub(BigNumber.from('1'))
  .toString();

export const ZERO_ADDRSSS = '0x0000000000000000000000000000000000000000';

export const ERC1155InterfaceId: string = '0xd9b67a26';

export const ERC721InterfaceId: string = '0x80ac58cd';

export const ChainConfigMap: {
  [key: string]: {
    chainId: string;
    chainName: string;
    rpcUrls: string[];
    blockExplorerUrls: string[];
    nativeCurrency: {
      name: string;
      symbol: string;
      decimals: number;
    };
    contractAddressMap: { [key: string]: string };
  };
} = {
  '0x5': {
    chainId: '0x5',
    chainName: 'Goerli',
    rpcUrls: ['https://rpc.ankr.com/eth_goerli'],
    blockExplorerUrls: ['https://goerli.etherscan.io'],
    nativeCurrency: {
      name: 'ETH',
      symbol: 'ETH',
      decimals: 18,
    },
    contractAddressMap: {
      FMToken: '0xDa741102500D71Fa1045C100415D0A5705d84eC5',
      ERC6551Account: '0x78d999776Eb78201deCa89D57f5aeAb7000e689f',
      AssistantFactory: '0xC61781ca60a105F23b3F8B52c81B70B4490cdbE2',
    },
  },
};

export const BackendHost = 'https://api.fusionmate.xyz';

export const AssistantTypeList: { name: string; value: AssistantType }[] = [
  {
    name: 'Web3 Expert',
    value: '0',
  },
  {
    name: 'Game Characters',
    value: '1',
  },
  {
    name: 'Novel Characters',
    value: '2',
  },
  {
    name: 'Celebrities',
    value: '3',
  },
  {
    name: 'Anime Girls',
    value: '4',
  },
];
