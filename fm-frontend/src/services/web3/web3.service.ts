import { Injectable } from '@angular/core';
import { ethers } from 'ethers';

import { BehaviorSubject, debounceTime, firstValueFrom, from, Observable, of, retry } from 'rxjs';
import type Web3 from "web3";
import type { TransactionReceipt, AbiParameter, AbiInput, FMT_NUMBER, FMT_BYTES } from "web3-types";
import { AccountService } from '../account/account.service';
import { WalletService } from '../wallet/wallet.service';

import { ChainInfoService } from '../chainInfo/chain-info.service';
import { ERC721InterfaceId, ERC1155InterfaceId } from 'src/common/config';
import { ContractReturnType, ContractSendType } from 'src/common/types';

@Injectable({
  providedIn: 'root'
})
export class Web3Service {

  private web3?: Web3;
  receipt$ = new BehaviorSubject<TransactionReceipt | null>(null);
  // recentTransaction: Set<{ txHash: string, contractName: string, functionName: string, params: Array<any>, receipt: any }> = new Set(JSON.parse(localStorage.getItem('recentTransaction') || '[]'));
  private abiMap: { [key: string]: { [key: string]: any } } = {};
  constructor(private walletService: WalletService, private accountService: AccountService, private chainInfoService: ChainInfoService) {
    this.initWeb3();
  }

  get ContractAddressMap() {
    return this.chainInfoService.getChainConfig().contractAddressMap;
  }

  initWeb3() {
    return new Promise<void>(async (resolve, reject) => {
      if (!this.web3) {
        Promise.all([import('web3'), this.walletService.getProvider()]).then(([Web3Module, provider]) => {
          this.web3 = new Web3Module.default(provider);
          resolve();
        }).catch(reject);
      } else if (!this.web3.currentProvider) {
        this.web3.setProvider(await this.walletService.getProvider());
        resolve();
      } else {
        resolve();
      }

    })

  }
  invokeContract<T extends ContractSendType>(contractName: string, functionName: string, params: Array<any>, address?: string, value?: string) {
    const contractAddress = address ?? this.ContractAddressMap[contractName];
    return this.initWeb3().then(() => {
      return Promise.all([this.getFunctionABI(contractName, functionName), firstValueFrom(this.accountService.account$)])
        .then(([functionAbi, currentAccount]): Promise<ContractReturnType<T>> => {
          const encodeData = this.web3!.eth.abi.encodeFunctionCall(functionAbi, params);
          const transactionObj = {
            from: currentAccount,
            to: contractAddress,
            data: encodeData,
            value
          };
          if (functionAbi.stateMutability == 'view') {
            return this.web3!.eth.call<{
              readonly number: FMT_NUMBER.STR;
              readonly bytes: FMT_BYTES.HEX;
            }>(transactionObj).then(
              returnValues => {
                return this.web3!.eth.abi.decodeParameters(functionAbi.outputs, returnValues) as ContractReturnType<T>
              }
            ).catch(
              e => {
                throw e
              }
            );
          } else {
            if (!this.walletService.isConnected) {
              throw new Error('Please connect wallet and switch to a Supported Blockchain.');
            }
            // let txhash = '';
            return this.web3!.eth.sendTransaction(transactionObj).on('transactionHash', (hash) => {
              // txhash = hash;
              // const arr = Array.from(this.recentTransaction);
              // arr.unshift({ txHash: hash, contractName, functionName, params, receipt: null });
              // this.recentTransaction = new Set(arr);
              // localStorage.setItem('recentTransaction', JSON.stringify(arr));
            })
              .on('receipt', (receipt) => {
                // this.recentTransaction.forEach((item) => {
                //   if (item.txHash === receipt.transactionHash) {
                //     item.receipt = receipt;
                //   }
                // });
                // localStorage.setItem('recentTransaction', JSON.stringify(Array.from(this.recentTransaction)));
                // if (receipt.status) {

                // } else {

                // }
              })
              // .on('confirmation', (confirmationNumber, receipt) => { })
              .on('error', ((err) => {

                // this.recentTransaction.forEach((item) => {
                //   if (item.txHash === txhash) {
                //     item.receipt = { status: false };
                //   }
                // });
                // localStorage.setItem('recentTransaction', JSON.stringify(Array.from(this.recentTransaction)));

              }))
              .then(receipt => {
                setTimeout(() => this.receipt$.next(receipt), 10000);
                this.accountService.account$.next(this.accountService.account$.value);
                return receipt as ContractReturnType<T>;
              });
          }
        });
    }).catch(
      e => {
        console.error('error:', contractName, functionName, params, contractAddress, value, e)
        throw e
      }
    );
  }
  sha3(msg: string) {
    return this.web3?.utils.sha3(msg);
  }
  decodeLog(inputs: AbiParameter[], hex: string, topics: string[]) {
    return this.web3?.eth.abi.decodeLog(inputs, hex, topics);
  }
  decodeParameters(abi: AbiInput[], bytes: string) {
    return this.web3?.eth.abi.decodeParameters(abi, bytes);
  }
  getFunctionABI(contractName: string, functionName: string) {
    let abi = this.abiMap[contractName] ? this.abiMap[contractName][functionName] : undefined;
    if (abi) {
      return Promise.resolve(abi);
    } else {
      return fetch(`/assets/abi/${contractName}.json`)
        .then(res => res.text())
        .then(abiJsonStr => JSON.parse(abiJsonStr).abi)
        .then((abi: Array<any>): Promise<any> => {
          const filterResult = abi.filter((item) => {
            return item.name == functionName && item.type == 'function';
          })
          if (filterResult.length === 0) {
            throw new Error('No such function:' + contractName + '-' + functionName);
          }
          this.abiMap[contractName] = {};
          this.abiMap[contractName][functionName] = filterResult[0];
          return filterResult[0];
        });
    }
  }
  getTokenDecimals(baseTokenAddress: string) {
    return this.invokeContract<'call'>('ERC20', 'decimals', [], baseTokenAddress);
  }
  getBaseTokenBalanceOf(address: string) {
    return this.initWeb3().then(
      () => this.web3!.eth.getBalance(address)
    )
  }
  getTokenName(contractAddress: string) {
    if (contractAddress.toLowerCase() == (this.ContractAddressMap['NativeToken'] as string).toLowerCase()) {
      return Promise.resolve({ 0: 'ETH' })
    }
    return this.initWeb3().then(
      () => firstValueFrom(this.accountService.account$).then(
        async currentAccount => {
          const transactionObj = {
            from: currentAccount,
            to: contractAddress,
            data: '0x06fdde03', //name()
          };
          const returnValues = await this.web3!.eth.call(transactionObj);
          return this.web3!.eth.abi.decodeParameters([
            {
              "internalType": "string",
              "name": "",
              "type": "string"
            }
          ], returnValues);
        }
      )
    )

  }

  getTokenSymbol(contractAddress: string) {
    if (contractAddress.toLowerCase() == (this.ContractAddressMap['NativeToken'] as string).toLowerCase()) {
      return Promise.resolve({ 0: 'ETH' });
    }
    return this.initWeb3().then(
      () => firstValueFrom(this.accountService.account$).then(
        async currentAccount => {
          const transactionObj = {
            from: currentAccount,
            to: contractAddress,
            data: '0x95d89b41', //symbol()
          };
          const returnValues = await this.web3!.eth.call(transactionObj);
          return this.web3!.eth.abi.decodeParameters([
            {
              "internalType": "string",
              "name": "",
              "type": "string"
            }
          ], returnValues);
        }
      )
    )
  }

  async isERC721(address: string): Promise<boolean> {
    try {
      const is721 = ((await this.invokeContract<'call'>('ERC165', 'supportsInterface', [ERC721InterfaceId], address)) as any)[0]
      return is721;
    } catch (error) {
      return false
    }
  }
  async isERC1155(address: string): Promise<boolean> {
    try {
      const is1155 = ((await this.invokeContract<'call'>('ERC165', 'supportsInterface', [ERC1155InterfaceId], address)) as any)[0]
      return is1155;
    } catch (error) {
      return false
    }
  }

  getUserNfts(collection: string, account?: string): Observable<{ ownedNfts: { tokenId: string }[]; totalCount: number; }> {
    // const nftList = new BehaviorSubject<{ ownedNfts: { title: string, tokenId: string, tokenUri: string }[]; totalCount: number; }>({ ownedNfts: [], totalCount: 0 });

    if (account) {
      return from(this.invokeContract<'call'>('Assistant', 'balanceOf', [account], collection).then(
        async res => {
          const total = Number(res[0]);
          const ownedNfts = await Promise.all(Array(total).fill(0).map((_, index) => {
            return this.invokeContract<'call'>('Assistant', 'tokenOfOwnerByIndex', [account, index + ''], collection).then(
              (result) => {
                const tokenId = (result[0] as BigInt).toString();
                return { tokenId };
              })
          }));
          return { ownedNfts, totalCount: total };
        }
      )
      )
    }
    return of({ ownedNfts: [], totalCount: 0 })
    // return nftList.pipe(debounceTime(500));
  }
}

