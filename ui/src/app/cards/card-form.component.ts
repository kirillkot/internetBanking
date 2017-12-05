import { Location } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { CURRENCIES } from '../const.module';
import { FormComponent } from '../abstract/form.component';

import { CardForm, Card, CardService } from './card.service';


@Component({
  selector: 'app-card-form',
  templateUrl: './card-form.component.html',
  styleUrls: ['./card-form.component.css']
})
export class CardFormComponent extends FormComponent<CardForm, Card>{
  currencies = CURRENCIES;

  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: CardService,
  ) {
    super(location, formbuilder, service);
  }

  fields(): any {
    return {
      name: ['', [Validators.required, Validators.minLength(4)]],
      currency: ['', [Validators.required, Validators.minLength(4)]],
      card_offer_id: [null, [Validators.required]],
    }
  }

}
