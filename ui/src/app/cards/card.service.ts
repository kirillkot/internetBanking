import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


export interface CardForm {
  name: string;
  currency: string;
  offer_id: number;
}

export interface Card {
  id: number;
  account_id: number;
  offer_id: number;

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

  block(id: number) {
    return this.http.put<Card>(`/api/cards/${id}/state/`, {});
  }

}
