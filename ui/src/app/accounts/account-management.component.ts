import { Component, ChangeDetectorRef } from '@angular/core';

import { ManageListComponent } from '../abstract/manage-list.component';

import { AccountForm, Account } from './account.service';
import { AccountService } from './account.service';


@Component({
  selector: 'app-account-management',
  templateUrl: './account-management.component.html',
  styleUrls: ['./account-management.component.css']
})
export class AccountManagementComponent extends
    ManageListComponent<AccountForm, Account> {
  displayedColumns = ['id', 'detail', 'currency', 'balance', 'add','move', 'actions'];

  constructor(
    service: AccountService,
    detector: ChangeDetectorRef,
  ) {
    super(service, detector);
  }

}
