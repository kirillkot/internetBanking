import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { CurrenciesManagementComponent } from './currencies-management.component';
import { CurrencyFormComponent } from './currency-form.component';

const routes: Routes = [
  { path: 'currencies/management', component: CurrenciesManagementComponent },
  { path: 'currencies/management/form', component: CurrencyFormComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class CurrenciesRoutingModule { }
