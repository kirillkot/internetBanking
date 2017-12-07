import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AccountManagementComponent } from '../accounts/account-management.component';


const routes: Routes = [
  { path: 'accounts/management', component: AccountManagementComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PaymentsRoutingModule { }
