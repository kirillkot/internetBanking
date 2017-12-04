import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { MaterialModule } from '../material/material.module';
import { CardsRoutingModule } from './cards-routing.module';

import { CardService } from './card.service';
import { CardOfferService } from '../card-offers/card-offer.service';

import { CardListComponent } from './card-list.component';
import { CardOfferListComponent } from '../card-offers/card-offer-list.component';
import { CardOfferFormComponent } from '../card-offers/card-offer-form.component';


@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
    MaterialModule,
    CardsRoutingModule
  ],
  declarations: [
    CardListComponent,
    CardOfferListComponent,
    CardOfferFormComponent,
  ],
  providers: [
    CardService,
    CardOfferService,
  ]
})
export class CardsModule { }
