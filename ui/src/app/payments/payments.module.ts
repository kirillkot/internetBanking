import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { MaterialModule } from '../material/material.module';
import { PaymentsRoutingModule } from './payments-routing.module';

import { AccountService } from '../accounts/account.service';
import { PaymentTypeService } from '../payment-types/payment-type.service';
import { PaymentService } from './payment.service';

import { AccountManagementComponent } from '../accounts/account-management.component';
import { AccountFormComponent } from '../accounts/account-form.component';
import { PaymentTypeManagementComponent } from '../payment-types/payment-type-management.component';
import { PaymentTypeFormComponent } from '../payment-types/payment-type-form.component';
import { PaymentManagementComponent } from './payment-management.component';


@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
    MaterialModule,
    PaymentsRoutingModule,
  ],
  declarations: [
    AccountManagementComponent,
    AccountFormComponent,
    PaymentTypeManagementComponent,
    PaymentTypeFormComponent,
    PaymentManagementComponent,
  ],
  providers: [
    AccountService,
    PaymentTypeService,
    PaymentService,
  ]
})
export class PaymentsModule { }
