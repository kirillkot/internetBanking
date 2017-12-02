import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { CardListComponent } from './card-list.component';
import { CardOfferListComponent } from '../card-offers/card-offer-list.component';

const routes: Routes = [
  { path: 'cards', component: CardListComponent },
  { path: 'cards/offers', component: CardOfferListComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class CardsRoutingModule { }
