import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


export interface TransactionForm {
}

export interface Transaction {
  id: number;
  account_id: number;

  delta: number;
  time: Date;
  detail: string;
}

@Injectable()
export class TransactionService extends
    BackendService<TransactionForm, Transaction> {

  constructor(
    http: HttpClient,
  ) {
    super('transactions', http);
  }

  getObjectsByCard(card: number) {
    return this.http.get< Transaction[] >(`/api/card/${card}/transactions`);
  }

}
