import { Location } from '@angular/common';
import { Observable } from 'rxjs/Observable';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { FormComponent } from '../abstract/form.component';

import { Card, CardService } from '../cards/card.service';
import { PaymentType, PaymentTypeService } from '../payment-types/payment-type.service';
import { PaymentForm, Payment, PaymentService } from './payment.service';


@Component({
  selector: 'app-payment-form',
  templateUrl: './payment-form.component.html',
  styleUrls: ['./payment-form.component.css']
})
export class PaymentFormComponent extends
    FormComponent<PaymentForm, Payment> implements
    OnInit {
  payment_types: Observable<PaymentType[]>;
  cards: Observable<Card[]>;

  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: PaymentService,
    private typeservice: PaymentTypeService,
    private cardservice: CardService,
  ) {
    super(location, formbuilder, service);
  }

  fields(): any {
    return {
      payment_type_id: [null, [Validators.required]],
      name: [null, [Validators.required]],
      from_account_id: [null, [Validators.required]],
      currency: [null, [Validators.required]],
      amount: [null, [Validators.required]],
    };
  }

  ngOnInit() {
    super.ngOnInit();
    this.payment_types = this.typeservice.getObjects();
    this.cards = this.cardservice.getObjects();
  }

  setPaymentType(type: PaymentType): void {
    this.group.patchValue({
      payment_type_id: type.id,
      name: type.name,
    });
  }

  setCard(card: Card): void {
    this.group.patchValue({
      from_account_id: card.account_id,
      currency: card.currency,
    });
  }

}
