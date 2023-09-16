import { NgFor, NgIf } from '@angular/common';
import { Component, Inject, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { RouterModule } from '@angular/router';

import { NzTimelineModule } from 'ng-zorro-antd/timeline';
import { NzSkeletonModule } from 'ng-zorro-antd/skeleton';

import { Assistant } from 'src/common/types';
import { AccountService } from 'src/services/account/account.service';
import { Web3Service } from 'src/services/web3/web3.service';

@Component({
  standalone: true,
  selector: 'app-mint-dialog',
  templateUrl: './mint-dialog.component.html',
  styleUrls: ['./mint-dialog.component.scss'],
  imports: [
    NgFor,
    NgIf,
    RouterModule,
    FormsModule,
    MatButtonModule,
    MatDialogModule,
    NzTimelineModule,
    NzSkeletonModule
  ],
})
export class MintDialogComponent implements OnInit {

  assistant: Assistant;
  pendingMsg = '';
  timelineList: string[] = ['Start'];
  isMinted = false;
  constructor(private dialogRef: MatDialogRef<MintDialogComponent>,
    @Inject(MAT_DIALOG_DATA) private data: Assistant, private web3Service: Web3Service, private accountService: AccountService) {
    this.assistant = data;
    this.dialogRef.disableClose = true;
  }
  ngOnInit(): void {
    this.approveToken().then(
      () => this.mint()
    )
  }

  approveToken() {
    // firstValueFrom(this.accountService.account$).then(
    //   account=>this.web3Service.invokeContract<'call'>('FMToken','allowance',[account,this.assistant.contractAddr])
    // )
    this.pendingMsg = 'Approving Token For the Tx...';
    return this.web3Service.invokeContract<'send'>('FMToken', 'approve', [this.assistant.contractAddr, this.assistant.mintPrice]).then(
      () => {
        this.pendingMsg = '';
        this.timelineList.push('FMToken Appproved');
      }
    ).catch(e => {
      this.pendingMsg = '';
      this.timelineList.push('Approve Token Failed.');
      this.dialogRef.disableClose = false;
    });
  }

  mint() {
    this.pendingMsg = 'Mint Assistant...';
    return this.web3Service.invokeContract<'send'>('Assistant', 'mint', [], this.assistant.contractAddr).then(
      () => {
        this.pendingMsg = '';
        this.timelineList.push('Assistant Minted.');
        this.isMinted = true;
        this.dialogRef.disableClose = false;
      }
    ).catch(e => {
      this.pendingMsg = '';
      this.timelineList.push('Mint Assistant Failed.');
      this.dialogRef.disableClose = false;
    });
  }


}
