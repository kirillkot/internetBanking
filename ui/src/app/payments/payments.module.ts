import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PaymentsRoutingModule } from './payments-routing.module';
import { AccountsModule } from '../accounts/accounts.module';

@NgModule({
  imports: [
    CommonModule,
    PaymentsRoutingModule,
    AccountsModule
  ],
  declarations: []
})
export class PaymentsModule { }
