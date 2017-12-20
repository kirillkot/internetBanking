import { Location } from '@angular/common';
import { Observable } from 'rxjs/Observable';
import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';

import { FormComponent } from '../abstract/form.component';

import { Card, CardService } from '../cards/card.service';
import { ConverterService } from '../converter.service';
import { PaymentType, PaymentTypeService } from '../payment-types/payment-type.service';
import { PaymentForm, Payment, PaymentService } from './payment.service';


@Component({
  selector: 'app-payment-form',
  templateUrl: './payment-form.component.html',
  styleUrls: ['./payment-form.component.css']
})
export class PaymentFormComponent extends
    FormComponent<PaymentForm, Payment> implements
    OnInit {
  payment_types: Observable<PaymentType[]>;
  cards: Observable<Card[]>;

  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: PaymentService,
    public dialog: MatDialog,
    private typeservice: PaymentTypeService,
    private cardservice: CardService,
  ) {
    super(location, formbuilder, service);
  }

  fields(): any {
    return {
      payment_type_id: [null, [Validators.required]],
      name: [null, [Validators.required]],
      from_account_id: [null, [Validators.required]],
      currency: [null, [Validators.required]],
      amount: [null, [Validators.required]],
      to_account_id: [null, [Validators.required]],
    };
  }

  ngOnInit() {
    super.ngOnInit();
    this.payment_types = this.typeservice.getObjects();
    this.cards = this.cardservice.getObjects();
  }

  setPaymentType(type: PaymentType): void {
    this.group.patchValue({
      payment_type_id: type.id,
      name: type.name,
      to_account_id: type.account_id,
    });
  }

  setCard(card: Card): void {
    this.group.patchValue({
      from_account_id: card.account_id,
      currency: card.currency,
    });
  }

  create(): void {
    let dialogRef = this.dialog.open(PaymentConverterComponent, {
      data: this.group.value,
    });

    dialogRef.afterClosed()
      .subscribe(
        (data) => {
          const flag = data as boolean;
          if (flag === true) {
            super.create();
          }
        },
      );
  }

}

@Component({
  selector: 'user-qr-component',
  template: `
  <h1 mat-dialog-title> Converter Dialog </h1>
  <div mat-dialog-content>
    <p>Do you want convert?</p>
    <p>{{ amount }} {{  amount_currency}}</p>
    <p>To</p>
    <p *ngIf="result"> {{ result }} {{ result_currency }}</p>
  </div>
  <div mat-dialog-actions>
    <button mat-button [mat-dialog-close]="true" tabindex="2">Ok</button>
    <button mat-button [mat-dialog-close]="false" tabindex="-1">No Thanks</button>
  </div>
`,
})
export class PaymentConverterComponent implements OnInit {
  public amount: string;
  public amount_currency: string;
  public result: string;
  public result_currency: string;

  constructor(
    public dialogRef: MatDialogRef<PaymentConverterComponent>,
    @Inject(MAT_DIALOG_DATA) private data: PaymentForm,
    private service: ConverterService,
  ) {
    this.amount = data.amount;
    this.amount_currency = data.currency;
  }

  ngOnInit(): void {
    this.service.convert({
      amount: this.data.amount,
      from: this.data.from_account_id,
      to: this.data.payment_type_id,
    }).subscribe(
      (data) => {
        this.amount = data.amount;
        this.amount_currency = data.amount_currency;
        this.result = data.result;
        this.result_currency = data.result_currency;
      },
    );
  }
}
