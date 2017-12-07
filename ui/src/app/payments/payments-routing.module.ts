import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AccountManagementComponent } from '../accounts/account-management.component';
import { PaymentTypeManagementComponent } from './payment-type-management.component';
import { PaymentTypeFormComponent } from './payment-type-form.component';


const routes: Routes = [
  { path: 'accounts/management', component: AccountManagementComponent },
  { path: 'payment-types/management', component: PaymentTypeManagementComponent },
  { path: 'payment-types/management/form', component: PaymentTypeFormComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PaymentsRoutingModule { }
