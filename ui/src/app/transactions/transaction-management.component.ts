import { Component, ChangeDetectorRef } from '@angular/core';

import { ManageListComponent } from '../abstract/manage-list.component';

import { TransactionForm, Transaction } from './transaction.service';
import { TransactionService } from './transaction.service';


@Component({
  selector: 'app-transaction-management',
  templateUrl: './transaction-management.component.html',
  styleUrls: ['./transaction-management.component.css']
})
export class TransactionManagementComponent extends
    ManageListComponent<TransactionForm, Transaction> {
  displayedColumns = ['id', 'account', 'delta', 'time', 'detail']

  constructor(
    service: TransactionService,
    detector: ChangeDetectorRef,
  ) {
    super(service, detector);
  }

}
