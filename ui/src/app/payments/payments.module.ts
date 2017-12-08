import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { MaterialModule } from '../material/material.module';
import { PaymentsRoutingModule } from './payments-routing.module';

import { AccountService } from '../accounts/account.service';
import { PaymentTypeService } from '../payment-types/payment-type.service';

import { AccountFormComponent } from '../accounts/account-form.component';
import { AccountManagementComponent } from '../accounts/account-management.component';
import { PaymentTypeManagementComponent } from '../payment-types/payment-type-management.component';
import { PaymentTypeFormComponent } from '../payment-types/payment-type-form.component';


@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
    MaterialModule,
    PaymentsRoutingModule,
  ],
  declarations: [
    AccountManagementComponent,
    PaymentTypeManagementComponent,
    PaymentTypeFormComponent,
    AccountFormComponent,
  ],
  providers: [
    AccountService,
    PaymentTypeService,
  ]
})
export class PaymentsModule { }
