import { Component, ChangeDetectorRef } from '@angular/core';

import { ManageListComponent } from '../abstract/manage-list.component';

import { PaymentForm, Payment } from './payment.service';
import { PaymentService } from './payment.service';


@Component({
  selector: 'app-payment-management',
  templateUrl: './payment-management.component.html',
  styleUrls: ['./payment-management.component.css']
})
export class PaymentManagementComponent extends
    ManageListComponent<PaymentForm, Payment> {
  displayedColumns = ['id', 'name', 'type', 'currency',
    'amount','commision', 'from', 'to', 'actions'];

  constructor(
    service: PaymentService,
    detector: ChangeDetectorRef,
  ) {
    super(service, detector);
  }

  repeat(payment: Payment) {
    return this.service.create(payment)
      .subscribe(
        (data) => this.refresh(),
      );
  }
}
