import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AccountManagementComponent } from '../accounts/account-management.component';
import { PaymentTypeManagementComponent } from './payment-type-management.component';


const routes: Routes = [
  { path: 'accounts/management', component: AccountManagementComponent },
  { path: 'payment-types/management', component: PaymentTypeManagementComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PaymentsRoutingModule { }
