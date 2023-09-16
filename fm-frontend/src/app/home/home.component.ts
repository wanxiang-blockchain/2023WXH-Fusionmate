import { Component } from '@angular/core';
import { BehaviorSubject, Observable, combineLatest, firstValueFrom, map, of, take } from 'rxjs';
// import { mockAssistant } from 'src/common/mock';
import { Assistant, AssistantType } from 'src/common/types';
import { CacheService } from 'src/services/cache/cache.service';
import { MatDialog } from '@angular/material/dialog';
import { MintDialogComponent } from '../mint-dialog/mint-dialog.component';
import { AccountService } from 'src/services/account/account.service';
import { WalletService } from 'src/services/wallet/wallet.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent {

  type$ = new BehaviorSubject<AssistantType | undefined>(undefined);
  assistantList$;
  listCount = 9;
  // mockAssistant = mockAssistant;
  constructor(private cacheService: CacheService, private matDialog: MatDialog, private accountService: AccountService, private walletService: WalletService) {
    this.assistantList$ = combineLatest([this.type$, this.cacheService.assistantList$]).pipe(
      map(([type, assistantList]) => {
        return assistantList.filter(assistant => type ? assistant.type == type : true)
      })
    ).pipe(take(this.listCount))

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

  filterAssistantList(type?: AssistantType) {
    this.type$.next(type);
  }
}
