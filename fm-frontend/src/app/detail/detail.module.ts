import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DetailRoutingModule } from './detail-routing.module';
import { DetailComponent } from './detail.component';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatDialogModule } from '@angular/material/dialog';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    DetailComponent
  ],
  imports: [
    CommonModule,
    DetailRoutingModule,
    FormsModule,
    MatButtonModule,
    MatIconModule,
    MatDialogModule
  ]
})
export class DetailModule { }
