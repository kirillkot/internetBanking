import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { MaterialModule } from '../material/material.module';
import { PaymentsRoutingModule } from './payments-routing.module';

import { AccountService } from '../accounts/account.service';
import { PaymentTypeService } from './payment-type.service';

import { AccountManagementComponent } from '../accounts/account-management.component';
import { PaymentTypeManagementComponent } from './payment-type-management.component';
import { PaymentTypeFormComponent } from './payment-type-form.component';


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
  ],
  providers: [
    AccountService,
    PaymentTypeService,
  ]
})
export class PaymentsModule { }
