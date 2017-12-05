import { Location } from '@angular/common';
import { Observable } from 'rxjs/Observable';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { CURRENCIES } from '../const.module';
import { FormComponent } from '../abstract/form.component';

import { CardForm, Card, CardService } from './card.service';
import { CardOffer, CardOfferService } from '../card-offers/card-offer.service';


@Component({
  selector: 'app-card-form',
  templateUrl: './card-form.component.html',
  styleUrls: ['./card-form.component.css']
})
export class CardFormComponent extends FormComponent<CardForm, Card> implements OnInit {
  currencies = CURRENCIES;
  offers: Observable<CardOffer[]>;

  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: CardService,
    private offerservice: CardOfferService,
  ) {
    super(location, formbuilder, service);
  }

  fields(): any {
    return {
      name: ['', [Validators.required]],
      currency: [this.currencies[0], [Validators.required]],
      offer_id: [null, [Validators.required]],
    };
  }

  ngOnInit() {
    super.ngOnInit();
    this.offers = this.offerservice.getObjects();
  }

  setOffer(offer: CardOffer): void {
    this.group.patchValue({offer_id: offer.id, name: offer.name});
  }

  create() {
    this.service.create(this.group.value);
  }

}
