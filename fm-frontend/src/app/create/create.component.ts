import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { AssistantTypeList } from 'src/common/config';
import { AssistantType } from 'src/common/types';
import { CacheService } from 'src/services/cache/cache.service';
import { DataService } from 'src/services/data/data.service';
import { CreateDialogComponent } from '../create-dialog/create-dialog.component';
import { firstValueFrom } from 'rxjs';
import { AccountService } from 'src/services/account/account.service';
import { WalletService } from 'src/services/wallet/wallet.service';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnInit {

  assistantList$;
  assistantTypeList = AssistantTypeList;
  name: string = '';
  description: string = '';
  replication: number = 0;
  mintPrice: number = 0;
  prompt: string = '';
  type: AssistantType = '0';
  derive?: string;
  constructor(private cacheService: CacheService, private dataService: DataService, private matDialog: MatDialog, private accountService: AccountService, private walletService: WalletService) {
    this.assistantList$ = this.cacheService.assistantList$;
  }
  ngOnInit(): void {
    firstValueFrom(this.accountService.account$).then(
      async account => {
        if (account) {
          if (!this.accountService.getToken()) {
            try {
              await this.walletService.signMessage(account, 'Hello FusionMate!');
            } catch (error) {
              return;
            }
          }
        } else {
          this.walletService.connect()
        }
      }
    );
  }
  promptInputChange(fileInputEvent: any) {
    const reader = new FileReader();
    reader.readAsText(fileInputEvent.target.files[0]);
    reader.onload = () => { this.prompt = reader.result as string };
  }
  create() {
    if (this.name && this.description && this.replication && this.replication > 0 && this.mintPrice && this.mintPrice > 0 && this.prompt && this.type) {
      firstValueFrom(this.accountService.account$).then(
        async account => {
          if (account) {
            if (!this.accountService.getToken()) {
              try {
                const signed = await this.walletService.signMessage(account, 'Hello FusionMate!');
                if (signed == false) {
                  setTimeout(() => this.create(), 500);
                  return;
                }
              } catch (error) {
                return;
              }
            }
            this.openCreateDialog();
          } else {
            this.walletService.connect().then(() => setTimeout(() => this.create(), 500));
          }
        }
      );

    }
  }

  openCreateDialog() {
    this.matDialog.open(CreateDialogComponent, {
      data: {
        name: this.name,
        description: this.description,
        replication: this.replication,
        mintPrice: this.mintPrice,
        prompt: this.prompt,
        type: this.type,
        derive: this.derive,
      }
    })
  }
}
