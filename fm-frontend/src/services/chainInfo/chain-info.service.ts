import { Injectable } from '@angular/core';
import { values } from 'lodash';
import { BehaviorSubject } from 'rxjs';
import { ChainConfigMap } from 'src/common/config';

@Injectable({
  providedIn: 'root'
})
export class ChainInfoService {

  prevChainId?: string;
  networks = values(ChainConfigMap);
  currentChainId: string | undefined;
  currentChainId$ = new BehaviorSubject<string | undefined>(undefined);
  constructor() { }
  setChainId(chainId: string) {
    if (chainId == this.currentChainId) {
      return;
    }
    this.currentChainId = chainId;
    this.currentChainId$.next(chainId);
  }
  getPrevChainId() {
    return this.prevChainId ?? localStorage.getItem('prevChainId') ?? this.networks[0].chainId
  }

  setPrevChainId(chainId: string) {
    this.prevChainId = chainId;
    localStorage.setItem('prevChainId', chainId);
  }
  getChainConfig() {
    return ChainConfigMap[this.getPrevChainId()];
  }
  isSupportedChain(chainId: any) {
    return !!this.networks.find(chain => chain.chainId == chainId);
  }
}
