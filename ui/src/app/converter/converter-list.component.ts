import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, Validators } from '@angular/forms';

import { ManageListComponent } from '../abstract/manage-list.component';

import { CURRENCIES } from '../const.module';
import { ConverterForm, Converter } from './converter.service';
import { ConverterService } from './converter.service';

@Component({
    selector: 'app-converter-list',
    templateUrl: './converter-list.component.html',
    styleUrls: ['./converter-list.component.css' ]
})
export class ConverterListComponent extends ManageListComponent<ConverterForm, Converter> {
    currencies: { id: number, name: string }[] = [
        { "id": 0, "name": "BYN" },
        { "id": 1, "name": "USD" },
        { "id": 2, "name": "RUB" },
        { "id": 3, "name": "EUR" }
    ];

    str: number;
    valueTo: number;
    fromCurr: string;

    constructor(
        detector: ChangeDetectorRef,
        private router: Router,
        private converterservice: ConverterService,
    ) {
        super(converterservice, detector);
    }

    fields(): any {
        return {
            currency: [this.currencies[0], [Validators.required]],
        };
    }

    calculate(): void {
        this.valueTo = this.str;
    }
}
