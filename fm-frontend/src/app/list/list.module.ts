import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ListRoutingModule } from './list-routing.module';
import { ListComponent } from './list.component';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogModule } from '@angular/material/dialog';

@NgModule({
  declarations: [
    ListComponent
  ],
  imports: [
    CommonModule,
    ListRoutingModule,
    MatButtonModule,
    MatDialogModule
  ]
})
export class ListModule { }
