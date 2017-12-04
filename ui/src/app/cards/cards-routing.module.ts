import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { CardListComponent } from './card-list.component';
import { CardOfferListComponent } from '../card-offers/card-offer-list.component';
import { CardOfferFormComponent } from '../card-offers/card-offer-form.component';

const routes: Routes = [
  { path: 'cards', component: CardListComponent },
  { path: 'card-offers', component: CardOfferListComponent },
  { path: 'card-offers/form', component: CardOfferFormComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class CardsRoutingModule { }
