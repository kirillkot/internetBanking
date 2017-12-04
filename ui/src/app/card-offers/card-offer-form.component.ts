import { Location } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { FormComponent } from '../abstract/form.component';

import { CardOfferForm, CardOffer, CardOfferService } from './card-offer.service';


@Component({
  selector: 'app-card-offer-form',
  templateUrl: './card-offer-form.component.html',
  styleUrls: ['./card-offer-form.component.css']
})
export class CardOfferFormComponent extends FormComponent<CardOfferForm, CardOffer> {
  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: CardOfferService,
  ) {
    super(location, formbuilder, service);
  }

  fields(): any {
    return {
      username: ['', [Validators.required, Validators.minLength(4)]],
      type: ['', [Validators.required, Validators.minLength(4)]],
      cashback: ['', [Validators.required]],
      currency: ['', [Validators.required]],
      ttlMonth: [12, [Validators.required]],
    };
  }

}
