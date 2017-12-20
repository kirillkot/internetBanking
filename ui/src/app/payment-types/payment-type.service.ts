import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


export interface PaymentTypeForm {
  name: string;

  type: string;
  commision: string;
  account_id: number;
  currency: string;

  detail: string;
}

export interface PaymentType extends PaymentTypeForm {
  id: number;
}

@Injectable()
export class PaymentTypeService extends
    BackendService<PaymentTypeForm, PaymentType> {
  constructor(
    http: HttpClient,
  ) {
    super('payment-types', http);
  }

}
