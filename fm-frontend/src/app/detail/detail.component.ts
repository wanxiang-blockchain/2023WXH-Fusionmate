import { Component, ElementRef, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Assistant } from 'src/common/types';
import { CacheService } from 'src/services/cache/cache.service';
import { MintDialogComponent } from '../mint-dialog/mint-dialog.component';
// import { mockAssistant } from 'src/common/mock';
import { Observable, Subscription, combineLatest, distinctUntilChanged, firstValueFrom, lastValueFrom, map, of, switchMap, tap, timer } from 'rxjs';
import { DataService } from 'src/services/data/data.service';
import { AccountService } from 'src/services/account/account.service';
import { WalletService } from 'src/services/wallet/wallet.service';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { Web3Service } from 'src/services/web3/web3.service';
import { parseUnits } from 'ethers/lib/utils';

@Component({
  selector: 'app-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.scss']
})
export class DetailComponent implements OnInit, OnDestroy {
  msgList: { msg$: Observable<string>, from: 0 | 1 }[] = [{
    msg$: of('Hi~'),
    from: 0
  }]

  collectionId?: string;
  assistant?: Assistant;

  msg = '';

  @ViewChild('talkField') private talkField?: ElementRef;
  assistantSubscription: Subscription;

  constructor(private cacheService: CacheService, private matDialog: MatDialog,
    private dataService: DataService, private accountService: AccountService, private walletService: WalletService, private route: ActivatedRoute, private web3Service: Web3Service) {
    this.assistantSubscription = this.route.paramMap.pipe(
      switchMap(
        (params: ParamMap) => {
          this.collectionId = params.get('collectionID') || undefined;
          if (this.collectionId) {
            return this.cacheService.getAssistantById(this.collectionId);
          } else {
            return of(undefined)
          }

        }
      )
    ).subscribe(
      async assistant => {
        this.assistant = assistant;
      }
    )

  }
  ngOnDestroy(): void {
    this.assistantSubscription.unsubscribe()
  }
  ngOnInit(): void {
  }
  typedMsg(msg$: Observable<string>) {
    return combineLatest([timer(800, 30), msg$]).pipe(map(([count, msg]) => msg.substring(0, count)))
  }
  harvest(e: Event, assistant: Assistant) {
    if (!this.assistant) {
      return;
    }
    e.stopPropagation();
    firstValueFrom(this.accountService.account$).then(
      async account => {
        if (account) {
          if (!this.accountService.getToken()) {
            try {
              const signed = await this.walletService.signMessage(account, 'Hello FusionMate!');
              if (signed == false) {
                setTimeout(() => this.harvest(e, assistant), 500);
                return;
              }
            } catch (error) {
              return;
            }
          }
          const resp = await firstValueFrom(this.assistant!.ownedTokenId$).then(
            tokenIds => {
              if (tokenIds.length > 0) {
                return firstValueFrom(this.dataService.harvest(assistant.collectionID, tokenIds[0]));
              }
              return
            }
          );
          if (resp) {
            await this.web3Service.invokeContract<'send'>('Assistant', 'harvestForTBA', [resp.data.signature, resp.data.tokenID + '', resp.data.tokenNum + '', resp.data.collectionID + ''], this.assistant!.contractAddr)
          }


        } else {
          this.walletService.connect().then(() => setTimeout(() => this.mint(e, assistant), 500));
        }
      }
    );
  }
  mint(e: Event, assistant: Assistant) {
    e.stopPropagation();
    firstValueFrom(this.accountService.account$).then(
      async account => {
        if (account) {
          if (!this.accountService.getToken()) {
            try {
              const signed = await this.walletService.signMessage(account, 'Hello FusionMate!');
              if (signed == false) {
                setTimeout(() => this.mint(e, assistant), 500);
                return;
              }
            } catch (error) {
              return;
            }
          }
          this.matDialog.open(MintDialogComponent, { data: assistant });
        } else {
          this.walletService.connect().then(() => setTimeout(() => this.mint(e, assistant), 500));
        }
      }
    );

  }

  sendMsg() {
    if (!this.msg) {
      return;
    }
    if (!this.assistant) {
      return;
    }
    firstValueFrom(this.assistant.ownedTokenId$).then(
      tokenIds => {
        if (tokenIds.length > 0) {
          this.userSpeak();
          const respMsg$ = this.dataService.sendMsg(this.assistant!.collectionID, tokenIds[0], this.msg);
          this.botSpeak(respMsg$);
          this.msg = '';
        }
      }
    )

  }

  userSpeak() {
    this.msgList.push({
      msg$: of(this.msg),
      from: 1
    });
    this.scrollToBottom();
  }

  botSpeak(msg$: Observable<string>) {
    this.msgList.push({
      msg$: this.typedMsg(msg$),
      from: 0
    });
    this.scrollToBottom();
  }

  scrollToBottom(): void {
    setTimeout(() => this.talkField!.nativeElement.scrollTop = this.talkField!.nativeElement.scrollHeight, 200);
  }
}
