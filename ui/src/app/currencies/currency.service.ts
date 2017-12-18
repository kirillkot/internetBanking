import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


export interface CurrencyForm {
  name: string;
  koef: number;
  sale: number;
  purchase: number;
}

export interface Currency {
  id: number;

  name: string;
  koef: number;
  sale: number;
  purchase: number;
}

@Injectable()
export class CurrencyService extends
    BackendService<CurrencyForm, Currency> {

  constructor(
    http: HttpClient,
  ) {
    super('currencies', http);
  }

}
