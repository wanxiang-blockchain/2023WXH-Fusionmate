import { Injectable } from '@angular/core';
import { parseUnits } from 'ethers/lib/utils';
import { BehaviorSubject, firstValueFrom, from, map, Observable, switchMap, timer, debounceTime, Subscription, debounce, distinctUntilChanged, of, retry, catchError, expand, EMPTY, concatMap, reduce, defer, mergeMap, concat, tap, filter } from 'rxjs';
import { AccountService } from '../account/account.service';
import { Web3Service } from '../web3/web3.service';
import { Assistant, AssistantType, FetchAIListResponse } from 'src/common/types';
import { DataService } from '../data/data.service';
import { toArray } from 'lodash';
import { ZERO_ADDRSSS } from 'src/common/config';
@Injectable({
  providedIn: 'root'
})
export class CacheService {


  // public timer$: Observable<void>;
  public assistantList$ = new BehaviorSubject<Assistant[]>([]);
  subFetchAssistantList?: Subscription;
  constructor(private web3Service: Web3Service, private accountService: AccountService, private dataService: DataService) {
    // this.timer$ = timer(0, 60000).pipe(map(
    //   () => {
    //     this.triggerRefresh()
    //   }
    // ));
    // timer(0, 60000).pipe(
    //   switchMap(() => this.fetchAssistantList().pipe(
    //     map(assistantList => {
    //       if (this.assistantList$.value.length <= assistantList.length) {
    //         this.assistantList$.next(assistantList);
    //       }
    //     }
    //     )))).subscribe();

    // this.timer$.subscribe();
    this.subFetchAssistantList = this.fetchAssistantList().pipe(
      map(assistantList => {
        if (this.assistantList$.value.length <= assistantList.length) {
          this.assistantList$.next(assistantList);
        }
      }
      )).subscribe();
  }

  triggerRefresh() {

  }

  fetchAssistantList(pageIndex: number = 0): Observable<Assistant[]> {
    let pageSize = 100;
    return this.accountService.account$.pipe(
      switchMap(
        account => defer(() => this.dataService.fetchAIList(pageIndex, pageSize)).pipe(
          mergeMap(({ data }) => {

            const assistantList: Assistant[] = data.filter(item => item.contractAddr != '' && item.contractAddr != ZERO_ADDRSSS).map(item => {
              const ownedNfts$ = this.web3Service.getUserNfts(item.contractAddr, account);
              return {
                ...item,
                mintAmount$: ownedNfts$.pipe(map(ownedNfts => ownedNfts.totalCount)),
                owned$: ownedNfts$.pipe(map(ownedNfts => ownedNfts.totalCount > 0)),
                isMaker$: this.accountService.account$.pipe(map(account => account?.toLowerCase() == item.maker.toLowerCase())),
                ownedTokenId$: ownedNfts$.pipe(map(ownedNfts => ownedNfts.ownedNfts.map(nft => nft.tokenId)))
              };
            })
            const items$ = of(assistantList);
            const next$: Observable<Assistant[]> = data.length == pageSize ? this.fetchAssistantList(++pageIndex) : EMPTY;
            return concat(items$, next$);
          })
        )
      )
    );
  }

  getAssistantById(collectionId: string) {
    return this.assistantList$.pipe(map(
      assistantList => assistantList.find(assistant => assistant.collectionID == collectionId)
    ));
  }


}
