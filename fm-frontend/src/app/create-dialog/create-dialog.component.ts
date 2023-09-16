import { NgFor, NgIf } from '@angular/common';
import { Component, Inject, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { RouterModule } from '@angular/router';
import { NzTimelineModule } from 'ng-zorro-antd/timeline';
import { NzSkeletonModule } from 'ng-zorro-antd/skeleton';

import { firstValueFrom } from 'rxjs';
import { AssistantType } from 'src/common/types';
import { DataService } from 'src/services/data/data.service';
import { Web3Service } from 'src/services/web3/web3.service';
import { parseUnits } from 'ethers/lib/utils';

@Component({
  standalone: true,
  selector: 'app-create-dialog',
  templateUrl: './create-dialog.component.html',
  styleUrls: ['./create-dialog.component.scss'],
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
export class CreateDialogComponent implements OnInit {

  pendingMsg = '';
  imgUrl = ''
  timelineList: string[] = ['Start'];
  collectionID = '';
  isCreated = false;
  constructor(private dialogRef: MatDialogRef<CreateDialogComponent>, @Inject(MAT_DIALOG_DATA) private data: {
    name: string;
    description: string;
    replication: number;
    mintPrice: number;
    prompt: string;
    type: AssistantType;
    derive?: string;
  }, private dataService: DataService, private web3Service: Web3Service) {

  }
  ngOnInit(): void {
    this.generateImg().then(
      () => this.createAssistant()
    )
  }

  generateImg() {
    this.pendingMsg = 'Generate Image...';
    return firstValueFrom(this.dataService.genImgURI(this.data.prompt)).then(
      ({ data }) => {
        this.imgUrl = data.imgURI;
        this.pendingMsg = '';
        this.timelineList.push('Image Generated');
      }
    ).catch(e => {
      this.pendingMsg = '';
      this.timelineList.push('Generate Image Failed.');
      this.dialogRef.disableClose = false;
    });
  }

  createAssistant() {
    this.pendingMsg = 'Create Assistant...';
    const {
      name,
      description,
      replication,
      mintPrice,
      prompt,
      type,
      derive } = this.data;
    return firstValueFrom(this.dataService.CreateAI(name,
      description,
      replication,
      parseUnits(mintPrice + '', '18').toString(),
      // mintPrice * 10 ** 18,
      prompt,
      this.imgUrl,
      type,
      derive)).then(({ data }) => {
        const {
          collectionID,
          name,
          symbol,
          baseURI,
          maxSupply,
          mintPrice,
          signature
        } = data;
        this.collectionID = collectionID;
        return this.web3Service.invokeContract<'send'>('AssistantFactory', 'createAssistant', [signature, name, symbol, baseURI, collectionID, maxSupply, mintPrice]);

      }
      ).then(
        transactionReceipt => {
          console.log(transactionReceipt);
          const log = transactionReceipt.logs.find((item: any) => item.address.toLowerCase() == (this.web3Service.ContractAddressMap['AssistantFactory'] as string).toLowerCase())
          if (!(log && log.topics && log.topics[0].toString().toLowerCase() == '0xcb781daaebdd4349dbc7e45f9497a25ab7cc731f5f2c74c6fa9100958683a607')) {
            this.pendingMsg = '';
            this.isCreated = false;
            this.timelineList.push('Created assistant failed.');
            return;
          }
          // const result = this.web3Service.decodeLog([
          //   {
          //     "indexed": true,
          //     "internalType": "address",
          //     "name": "maker",
          //     "type": "address"
          //   },
          //   {
          //     "indexed": true,
          //     "internalType": "address",
          //     "name": "astBot",
          //     "type": "address"
          //   },
          //   {
          //     "indexed": true,
          //     "internalType": "uint256",
          //     "name": "nftId",
          //     "type": "uint256"
          //   }
          // ], log.data.toString(),
          //   ['0xcb781daaebdd4349dbc7e45f9497a25ab7cc731f5f2c74c6fa9100958683a607']);
          // if (!result) {
          //   this.pendingMsg = '';
          //   this.isCreated = false;
          //   this.timelineList.push('Created assistant failed.');
          //   return;
          // }
          const contractAddress = this.web3Service.decodeParameters(['address'], log.topics[2].toString())![0] as string;
          return firstValueFrom(this.dataService.notifyCreateResult(this.collectionID, contractAddress, true)).then(
            () => {
              this.pendingMsg = '';
              this.isCreated = true;
              this.timelineList.push('Assistant created');
            },
            () => {
              this.pendingMsg = '';
              this.isCreated = false;
              this.timelineList.push('Created assistant failed.');
            }
          );

        },
        () => {
          this.pendingMsg = '';
          this.isCreated = false;
          this.timelineList.push('Created assistant failed.');
        }
      )
  }



}
