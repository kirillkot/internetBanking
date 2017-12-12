import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';

export interface PaymentForm {
}

export interface Payment {
  id: number;

  name: string;
  type: string;

  currency: string;
  amount: number;
  commision: number;

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
