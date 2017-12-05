import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { CardListComponent } from './card-list.component';
import { CardOfferFormComponent } from '../card-offers/card-offer-form.component';
import { CardOfferManagementComponent } from '../card-offers/card-offer-management.component';


const routes: Routes = [
  { path: 'cards', component: CardListComponent },
  { path: 'card-offers/management', component: CardOfferManagementComponent },
  { path: 'card-offers/management/form', component: CardOfferFormComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class CardsRoutingModule { }
