import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


export interface AccountForm {
  currency: string;
  balance: number;

  add_allow: boolean;
  move_allow: boolean;

  detail: string;
}

export interface Account {
  id: number;

  currency: string;
  balance: number;

  add_allow: boolean;
  move_allow: boolean;

  detail: string;
}

@Injectable()
export class AccountService extends BackendService<AccountForm, Account> {
  constructor(
    http: HttpClient,
  ) {
    super('accounts', http);
  }

}
