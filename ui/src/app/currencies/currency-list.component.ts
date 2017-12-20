import { Component, ChangeDetectorRef } from '@angular/core';

import { ManageListComponent } from '../abstract/manage-list.component';

import { CurrencyForm, Currency } from './currency.service';
import { CurrencyService } from './currency.service';


@Component({
  selector: 'app-currency-list',
  templateUrl: './currency-list.component.html',
  styleUrls: ['./currency-list.component.css']
})
export class CurrencyListComponent extends
    ManageListComponent<CurrencyForm, Currency> {
  displayedColumns = ['name', 'sale', 'purchase'];

  constructor(
    service: CurrencyService,
    detector: ChangeDetectorRef,
  ) {
    super(service, detector);
  }

}
