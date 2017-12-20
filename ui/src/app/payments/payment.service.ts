import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


export interface PaymentForm {
  payment_type_id: number;
  payment_type_currency: string;
  name: string;
  from_account_id: number;
  currency: string;
  amount: string;
}

export interface Payment {
  id: number;
  payment_type_id: number;
  payment_type_currency: string;
  from_account_id: number;

  name: string;
  type: string;

  currency: string;
  amount: string;
  commision: string;

  from: string;
  to: string;
}


@Injectable()
export class PaymentService extends
    BackendService<PaymentForm, Payment> {

  constructor(
    http: HttpClient,
  ) {
    super('payments', http);
  }

}
