import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CreateRoutingModule } from './create-routing.module';
import { CreateComponent } from './create.component';
import { MatButtonModule } from '@angular/material/button';

import { NzSelectModule } from 'ng-zorro-antd/select';
import { FormsModule } from '@angular/forms';
import { MatIconModule } from '@angular/material/icon';
import { MatDialogModule } from '@angular/material/dialog';

@NgModule({
  declarations: [
    CreateComponent
  ],
  imports: [
    CommonModule,
    CreateRoutingModule,
    FormsModule,
    MatButtonModule,
    NzSelectModule,
    MatIconModule,
    MatDialogModule
  ]
})
export class CreateModule { }
