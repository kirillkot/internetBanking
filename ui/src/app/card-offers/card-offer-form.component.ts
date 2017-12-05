import { Location } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { CURRENCIES, CARD_TYPES } from '../const.module';
import { FormComponent } from '../abstract/form.component';

import { CardOfferForm, CardOffer, CardOfferService } from './card-offer.service';


@Component({
  selector: 'app-card-offer-form',
  templateUrl: './card-offer-form.component.html',
  styleUrls: ['./card-offer-form.component.css']
})
export class CardOfferFormComponent extends FormComponent<CardOfferForm, CardOffer> {
  cardTypes = CARD_TYPES;
  currencies = CURRENCIES;

  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: CardOfferService,
  ) {
    super(location, formbuilder, service);
  }

  fields(): any {
    return {
      name: ['', [Validators.required, Validators.minLength(4)]],
      type: [this.cardTypes[0], [Validators.required, Validators.minLength(4)]],
      details: [''],

      cashback: [1, [Validators.required]],
      currency: [this.currencies[0], [Validators.required]],
      ttlMonth: [12, [Validators.required]],
    };
  }

}
