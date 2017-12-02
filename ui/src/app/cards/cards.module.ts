import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { MaterialModule } from '../material/material.module';

import { CardsRoutingModule } from './cards-routing.module';
import { CardListComponent } from './card-list.component';

@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
    MaterialModule,
    CardsRoutingModule
  ],
  declarations: [
    CardListComponent,
  ]
})
export class CardsModule { }
