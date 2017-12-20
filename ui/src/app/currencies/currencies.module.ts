import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { MaterialModule } from '../material/material.module';

import { CurrenciesRoutingModule } from './currencies-routing.module';
import { CurrenciesManagementComponent } from './currencies-management.component';
import { CurrencyService } from './currency.service';
import { CurrencyFormComponent } from './currency-form.component';
import { CurrencyListComponent } from './currency-list.component';

@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
    MaterialModule,
    CurrenciesRoutingModule
  ],
  declarations: [CurrenciesManagementComponent, CurrencyFormComponent, CurrencyListComponent],
  providers: [CurrencyService]
})
export class CurrenciesModule { }
