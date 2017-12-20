import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


export interface AccountForm {
  currency: string;
  balance: string;

  add_allow: boolean;
  move_allow: boolean;

  detail: string;
}

export interface Account {
  id: number;
  iban: string;

  currency: string;
  balance: string;

  add_allow: boolean;
  move_allow: boolean;

  detail: string;
}

export interface AddFundsRequest {
  id: number;
  amount: number;
}

@Injectable()
export class AccountService extends BackendService<AccountForm, Account> {
  constructor(
    http: HttpClient,
  ) {
    super('accounts', http);
  }

  addFunds(req: AddFundsRequest) {
    return this.http.post(`/api/accounts/${req.id}/add/`,
      {amount: req.amount});
  }

}
