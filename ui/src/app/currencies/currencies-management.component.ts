import { Component, ChangeDetectorRef } from '@angular/core';

import { ManageListComponent } from '../abstract/manage-list.component';

import { CurrencyForm, Currency } from './currency.service';
import { CurrencyService } from './currency.service';


@Component({
  selector: 'app-currencies-management',
  templateUrl: './currencies-management.component.html',
  styleUrls: ['./currencies-management.component.css']
})
export class CurrenciesManagementComponent extends
    ManageListComponent<CurrencyForm, Currency> {
  displayedColumns = ['id', 'name', 'sale', 'purchase', 'actions'];

  constructor(
    service: CurrencyService,
    detector: ChangeDetectorRef,
  ) {
    super(service, detector);
  }

}
