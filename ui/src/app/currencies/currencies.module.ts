import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { MaterialModule } from '../material/material.module';

import { CurrenciesRoutingModule } from './currencies-routing.module';
import { CurrenciesManagementComponent } from './currencies-management.component';
import { CurrencyService } from './currency.service';

@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
    MaterialModule,
    CurrenciesRoutingModule
  ],
  declarations: [CurrenciesManagementComponent],
  providers: [CurrencyService]
})
export class CurrenciesModule { }
