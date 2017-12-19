import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AccountManagementComponent } from '../accounts/account-management.component';
import { AccountFormComponent } from '../accounts/account-form.component';
import { AccountDetailsComponent } from '../accounts/account-details.component';
import { PaymentTypeManagementComponent } from '../payment-types/payment-type-management.component';
import { PaymentTypeFormComponent } from '../payment-types/payment-type-form.component';
import { PaymentManagementComponent } from './payment-management.component';
import { PaymentFormComponent } from './payment-form.component';
import { TransactionListComponent } from '../transactions/transaction-list.component';
import { TransactionManagementComponent } from '../transactions/transaction-management.component';


const routes: Routes = [
  { path: 'accounts/management', component: AccountManagementComponent },
  { path: 'accounts/management/form', component: AccountFormComponent },
  { path: 'accounts/management/details', component: AccountDetailsComponent },
  { path: 'payment-types/management', component: PaymentTypeManagementComponent },
  { path: 'payment-types/management/form', component: PaymentTypeFormComponent },
  { path: 'payments/management', component: PaymentManagementComponent },
  { path: 'payments/management/form', component: PaymentFormComponent },
  { path: 'transactions', component: TransactionListComponent },
  { path: 'transactions/management', component: TransactionManagementComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PaymentsRoutingModule { }
