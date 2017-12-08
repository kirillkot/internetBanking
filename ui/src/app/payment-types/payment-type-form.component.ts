import { Location } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { FormComponent } from '../abstract/form.component';

import { PaymentTypeForm, PaymentType } from './payment-type.service';
import { PaymentTypeService } from './payment-type.service';


@Component({
  selector: 'app-payment-type-form',
  templateUrl: './payment-type-form.component.html',
  styleUrls: ['./payment-type-form.component.css']
})
export class PaymentTypeFormComponent extends
    FormComponent<PaymentTypeForm, PaymentType> {
  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: PaymentTypeService,
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

}
