import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


export interface CardOfferForm {
  name: string;
  type: string;
  cashback: string;
  currency: string;
  ttlMonth: number;
}

export interface CardOffer {
  id: number;
  name: string;
  type: string;
  cashback: string;
  currency: string;
  ttlMonth: number;
}

@Injectable()
export class CardOfferService extends BackendService<CardOfferForm, CardOffer> {
  cardTypes = ['Master Card', 'Visa', 'Visa-Electron', 'BelCard']

  constructor(
    http: HttpClient,
  ) {
    super('card-offers', http);
  }

}
