import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { MaterialModule } from '../material/material.module';
import { PaymentsRoutingModule } from './payments-routing.module';

import { AccountService } from '../accounts/account.service';
import { PaymentTypeService } from '../payment-types/payment-type.service';
import { PaymentService } from './payment.service';
import { TransactionService } from '../transactions/transaction.service';

import { AccountManagementComponent } from '../accounts/account-management.component';
import { AccountFormComponent } from '../accounts/account-form.component';
import { PaymentTypeManagementComponent } from '../payment-types/payment-type-management.component';
import { PaymentTypeFormComponent } from '../payment-types/payment-type-form.component';
import { PaymentManagementComponent } from './payment-management.component';
import { PaymentFormComponent } from './payment-form.component';
import { TransactionManagementComponent } from '../transactions/transaction-management.component';
import { AccountDetailsComponent } from '../accounts/account-details.component';
import { CurrenciesModule } from '../currencies/currencies.module';
import { TransactionListComponent } from '../transactions/transaction-list.component';


@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
    MaterialModule,
    CurrenciesModule,
    PaymentsRoutingModule,
  ],
  declarations: [
    AccountManagementComponent,
    AccountFormComponent,
    AccountDetailsComponent,
    PaymentTypeManagementComponent,
    PaymentTypeFormComponent,
    PaymentManagementComponent,
    PaymentFormComponent,
    TransactionManagementComponent,
    TransactionListComponent,
  ],
  providers: [
    AccountService,
    PaymentTypeService,
    PaymentService,
    TransactionService,
  ]
})
export class PaymentsModule { }
