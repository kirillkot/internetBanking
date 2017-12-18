import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CurrenciesRoutingModule } from './currencies-routing.module';
import { CurrenciesManagementComponent } from './currencies-management.component';

@NgModule({
  imports: [
    CommonModule,
    CurrenciesRoutingModule
  ],
  declarations: [CurrenciesManagementComponent]
})
export class CurrenciesModule { }
