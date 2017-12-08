import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


export interface PaymentTypeForm {
  name: string;

  type: string;
  commision: number;
  account_id: number;

  detail: string;
}

export interface PaymentType {
  id: number;

  name: string;

  type: string;
  commision: number;
  account_id: number;

  detail: string;
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
