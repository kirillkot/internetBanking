import { Location } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { CURRENCIES } from '../const.module';
import { FormComponent } from '../abstract/form.component';

import { AccountForm, Account, AccountService } from './account.service';


@Component({
  selector: 'app-account-form',
  templateUrl: './account-form.component.html',
  styleUrls: ['./account-form.component.css']
})
export class AccountFormComponent extends FormComponent<AccountForm, Account> {
  currencies = CURRENCIES;

  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: AccountService,
  ) {
    super(location, formbuilder, service);
  }

  fields(): any {
    return {
      currency: [this.currencies[0], [Validators.required]],
      balance: [0, [Validators.required]],

      add_allow: [true, [Validators.required]],
      move_allow: [true, [Validators.required]],
 
      details: [''],
    };
  }

}
