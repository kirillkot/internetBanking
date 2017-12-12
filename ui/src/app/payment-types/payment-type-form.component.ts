import { Location } from '@angular/common';
import { Observable } from 'rxjs/Observable';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { FormComponent } from '../abstract/form.component';

import { Account, AccountService } from '../accounts/account.service';
import { PaymentTypeForm, PaymentType } from './payment-type.service';
import { PaymentTypeService } from './payment-type.service';


@Component({
  selector: 'app-payment-type-form',
  templateUrl: './payment-type-form.component.html',
  styleUrls: ['./payment-type-form.component.css']
})
export class PaymentTypeFormComponent extends
    FormComponent<PaymentTypeForm, PaymentType> implements OnInit {
  accounts: Observable<Account[]>;

  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: PaymentTypeService,
    private accountservice: AccountService,
  ) {
    super(location, formbuilder, service);
  }

  fields(): any {
    return {
      name: ['', [Validators.required, Validators.minLength(4)]],
      type: ['', [Validators.required, Validators.minLength(2)]],
      commision: [0, [Validators.required]],
      account_id: [0, []],
      detail: ['', []],
    };
  }

  ngOnInit() {
    super.ngOnInit();
    this.accounts = this.accountservice.getObjects();
  }

  setAccount(account: Account): void {
    this.group.patchValue({account_id: account.id});
  }

}
