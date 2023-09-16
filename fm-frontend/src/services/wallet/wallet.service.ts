import { Injectable } from '@angular/core';
import { AccountService } from '../account/account.service';
import Onboard, { ConnectOptions, EIP1193Provider, OnboardAPI } from '@web3-onboard/core'
import injectedModule from '@web3-onboard/injected-wallets';
import { ChainInfoService } from '../chainInfo/chain-info.service';
import { BackendHost, ChainConfigMap } from 'src/common/config';
import { sign } from 'src/common/utils';
import { ethers } from 'ethers';



@Injectable({
  providedIn: 'root'
})
export class WalletService {

  isConnected = false;
  onboard: OnboardAPI;
  provider?: EIP1193Provider;
  isSigning = false;
  constructor(private accountService: AccountService, private chainInfoService: ChainInfoService) {
    this.onboard = Onboard({
      theme: {
        '--w3o-background-color': 'rgba(20,20,20)',
        '--w3o-foreground-color': 'rgba(50,50,50)',
        '--w3o-text-color': '#fff',
        '--w3o-border-color': 'rgba(50,50,50)',
        '--w3o-action-color': 'rgba(20,20,20)',
        '--w3o-border-radius': '1px',
      },
      wallets: [
        injectedModule()
      ],
      appMetadata: {
        name: 'FusionMate',
        description: 'FusionMate'
      },
      chains: this.chainInfoService.networks.map(
        network => ({
          id: network.chainId,
          token: network.nativeCurrency.symbol,
          label: network.chainName,
          rpcUrl: network.rpcUrls[0]
        })
      ),
      accountCenter: {
        desktop: {
          enabled: false
        },
        mobile: {
          enabled: false
        }
      },
      notify: {
        enabled: false
      },
      connect: {
        showSidebar: false
      }
    })
  }

  connect() {
    if (this.isConnected) {
      return Promise.resolve(true);
    }
    const connectedWallets = sessionStorage.getItem('connectedWallets')
    const previouslyConnectedWallets = connectedWallets ? JSON.parse(
      connectedWallets
    ) : undefined
    let option;
    if (previouslyConnectedWallets) {
      option = {
        autoSelect: { label: previouslyConnectedWallets[0], disableModals: true }
      }
    }
    return this.onboard.connectWallet(option).then(
      wallets => {
        if (wallets[0]) {
          this.provider = wallets[0].provider;

          // Subscribe to accounts change
          this.provider.on("accountsChanged", (accounts: string[]) => {
            setTimeout(async () => {
              if (accounts.length > 0) {
                this.accountService.account$.next(accounts[0].toLowerCase());
                if (!this.isSigning) {
                  await this.signMessage(accounts[0], 'Hello FusionMate!');
                }
                this.isConnected = accounts.length > 0;
                // this.refreshService.invokeFunctions();
              } else {
                this.isConnected = false;
                this.accountService.account$.next(undefined);
                this.accountService.setToken('');
              }
            });


          });

          // Subscribe to chainId change
          this.provider.on("chainChanged", (chainId) => {
            // console.log(chainId);
            this.chainInfoService.setChainId(chainId);

            if (!this.chainInfoService.networks.find(network => network.chainId == chainId)) {
              this.isConnected = false;
            } else {
              this.chainInfoService.setPrevChainId(chainId);
            }

          });

          // Subscribe to provider connection
          this.provider.on("connect", (info) => {
            this.isConnected = true;
            // console.log(info);
          });

          // Subscribe to provider disconnection
          this.provider.on("disconnect", (error) => {
            this.isConnected = false;
            this.accountService.account$.next(undefined);
            this.accountService.setToken('');
            // this.disconnect();
            // console.log(error);
          });
          // this.isConnected = this.networkIds.includes(provider.chainId);
          // if (!this.isConnected) {

          // }
          const prevChainId = this.chainInfoService.getPrevChainId();
          return this.provider?.request({ method: 'wallet_switchEthereumChain', params: [{ chainId: prevChainId }] })
            .then(() => {
              this.chainInfoService.setChainId(prevChainId);
              return this.provider!.request({ method: 'eth_accounts' })
                .then((accounts: string[]) => {
                  // console.log(accounts);
                  if (accounts.length > 0) {
                    this.isConnected = true;
                    this.accountService.account$.next(accounts[0].toLowerCase());
                    if (!this.isSigning) {
                      this.signMessage(accounts[0], 'Hello FusionMate!');
                    }
                    return true;
                  }
                  return false;
                });
            }, (e) => {
              console.error(e);
              if (e.code == 4902) {
                const {
                  chainId,
                  chainName,
                  rpcUrls,
                  blockExplorerUrls,
                  nativeCurrency
                } = ChainConfigMap[prevChainId];
                return this.provider?.request({
                  method: 'wallet_addEthereumChain',
                  params: [
                    {
                      chainId,
                      chainName,
                      rpcUrls,
                      blockExplorerUrls,
                      nativeCurrency
                    },
                  ],
                }).then(
                  () => this.provider?.request({ method: 'wallet_switchEthereumChain', params: [{ chainId: prevChainId }] })
                ).then(
                  () => {
                    this.isConnected = true;
                    this.provider!.request({ method: 'eth_chainId' }).then(
                      chainId => this.chainInfoService.setChainId(chainId)
                    )
                    return true
                  },
                  () => {
                    this.isConnected = false;
                    this.provider!.request({ method: 'eth_chainId' }).then(
                      chainId => this.chainInfoService.setChainId(chainId)
                    )
                    return false
                  }
                )
              } else {
                this.isConnected = false;
                this.provider!.request({ method: 'eth_chainId' }).then(
                  chainId => this.chainInfoService.setChainId(chainId)
                )
                return false
              }

            })
        }
        return false;
      }
    ).then(
      success => {
        const wallets = this.onboard.state.get().wallets;
        const connectedWallets = wallets.map(({ label }) => label)
        sessionStorage.setItem(
          'connectedWallets',
          JSON.stringify(connectedWallets)
        )
        return success;
      }
    )
  }

  async disconnect() {
    this.isConnected = false;
    this.accountService.account$.next(undefined);
    this.provider = undefined;
    const [primaryWallet] = this.onboard.state.get().wallets
    if (primaryWallet) {
      await this.onboard.disconnectWallet({ label: primaryWallet.label });
    }
    sessionStorage.removeItem('connectedWallets')
  }

  switchChain(chainId: string) {
    const chainConfig = ChainConfigMap[chainId];
    if (!this.provider) {
      return Promise.reject()
    }
    return this.provider.request({ method: 'wallet_switchEthereumChain', params: [{ chainId }] }).then(
      () => {
        // console.log('switchChain',this.provider)
        this.chainInfoService.setPrevChainId(chainId);
        return this.provider!.request({ method: 'eth_accounts' })
          .then((accounts: string[]) => {
            // console.log(accounts);
            if (accounts.length > 0) {
              this.isConnected = true;
              this.accountService.account$.next(accounts[0].toLowerCase());
              return true;
            }
            return false;
          });
      }
      , (e) => {
        if (e.code == 4902) {
          const {
            chainId,
            chainName,
            rpcUrls,
            blockExplorerUrls,
            nativeCurrency
          } = chainConfig;
          return this.provider?.request({
            method: 'wallet_addEthereumChain',
            params: [
              {
                chainId,
                chainName,
                rpcUrls,
                blockExplorerUrls,
                nativeCurrency
              },
            ],
          }).then(
            () => this.provider?.request({ method: 'wallet_switchEthereumChain', params: [{ chainId }] }).then(
              () => {
                this.chainInfoService.setPrevChainId(chainId);
                return this.provider!.request({ method: 'eth_accounts' })
                  .then((accounts: string[]) => {
                    // console.log(accounts);
                    if (accounts.length > 0) {
                      this.isConnected = true;
                      this.accountService.account$.next(accounts[0].toLowerCase());
                      return true;
                    }
                    return false;
                  });
              }
            )
          )
        }
        return;
      }
    )
  }



  getProvider() {
    return new Promise<any>((resolve, reject) => {
      if (this.provider) {
        resolve(this.provider);
      } else {
        const prevChainId = this.chainInfoService.getPrevChainId();
        this.connect().then(
          (isConnected) => {
            resolve(isConnected ? this.provider : ChainConfigMap[prevChainId].rpcUrls[0])
          }, () => resolve(ChainConfigMap[prevChainId].rpcUrls[0]));
      }
    });
  }

  async signMessage(address: string, msg: string) {
    if (this.isSigning) {
      return false;
    }
    if (this.accountService.getToken(address)) {
      return;
    }
    this.isSigning = true;
    const api = `${BackendHost}/api/v1`;
    const web3Provider = new ethers.providers.Web3Provider((window as any).ethereum);
    const domain = {
      name: 'FusionMate',
      version: '1',
      chainId: 5,
    }

    const primaryType = 'FusionMateChallenge'

    const types = {

      FusionMateChallenge: [
        { name: 'address', type: 'address' },
        { name: 'message', type: 'string' },
        { name: 'timestamp', type: 'string' }
      ],
      EIP712Domain: [
        { name: "name", type: "string" },
        { name: "version", type: "string" },
        { name: "chainId", type: "uint256" },
      ]
    }
    const signer = await web3Provider.getSigner();
    const timestamp = Math.floor(Date.now() / 1000).toString();

    const message = {
      "timestamp": timestamp,
      "address": address,
      "message": msg,
    }

    return new Promise((resolve, reject) => {
      // signer._signTypedData(domain, types, message).then(
      sign(domain, primaryType, message, types, signer).then(
        (result) => {
          let data = {
            "address": address,
            "message": msg,
            "timestamp": timestamp,
            "signature": result,
          };

          fetch(api + "/auth/login", {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
          }).then(res => res.json()).then((r) => {
            let returnCode = r.returnCode;
            //check return code here
            if (returnCode === 20000) {
              const { token } = r.data;
              if (token.length > 0) this.accountService.setToken(token);
              resolve(true);
            }
            reject({ message: "server error, please try again" });

          }).catch(error => {
            reject({ message: error })
          });
        }).catch(e => {
          this.isSigning = false;
          reject(e);
        });
    })
  }
}
