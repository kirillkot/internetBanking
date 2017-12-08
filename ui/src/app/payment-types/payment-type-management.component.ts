import { Component, ChangeDetectorRef } from '@angular/core';

import { ManageListComponent } from '../abstract/manage-list.component';

import { PaymentTypeForm, PaymentType } from './payment-type.service';
import { PaymentTypeService } from './payment-type.service';


@Component({
  selector: 'app-payment-type-management',
  templateUrl: './payment-type-management.component.html',
  styleUrls: ['./payment-type-management.component.css']
})
export class PaymentTypeManagementComponent extends
    ManageListComponent<PaymentTypeForm, PaymentType> {
  displayedColumns = ['id', 'name', 'type', 'commision',
      'account', 'detail', 'actions'];

  constructor(
    service: PaymentTypeService,
    detector: ChangeDetectorRef,
  ) {
    super(service, detector);
  }

}
