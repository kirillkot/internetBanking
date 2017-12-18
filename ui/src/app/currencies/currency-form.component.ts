import { Location } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { FormComponent } from '../abstract/form.component';

import { CurrencyForm, Currency, CurrencyService } from './currency.service';


@Component({
  selector: 'app-currency-form',
  templateUrl: './currency-form.component.html',
  styleUrls: ['./currency-form.component.css']
})
export class CurrencyFormComponent extends
    FormComponent<CurrencyForm, Currency> {

  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: CurrencyService,
  ) {
    super(location, formbuilder, service);
  }

  fields(): any {
    return {
      name: ['', [Validators.required]],
      koef: [1, [Validators.required]],
      sale: [null, [Validators.required]],
      purchase: [null, [Validators.required]],
    }
  }

}
