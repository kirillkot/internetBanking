import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';

export interface ConverterForm {
    currency: string;
}

export interface Converter {
    id: number;
}

@Injectable()
export class ConverterService extends BackendService<ConverterForm, Converter> {
    constructor(
        http: HttpClient,
    ){
        super('converter', http);
    }
}
