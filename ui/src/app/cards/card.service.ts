import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


export interface CardForm {
  name: string;
  currency: string;
  card_offer_id: number;
}

export interface Card {
  name: string;
  type: string;

  start_time: Date;
  valid_until: Date;
  status: string;

  currency: string;
  balance: string;
}

@Injectable()
export class CardService extends BackendService<CardForm, Card> {
  constructor(
    http: HttpClient,
  ) {
    super('cards', http);
  }

}
