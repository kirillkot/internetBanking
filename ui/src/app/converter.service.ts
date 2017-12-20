import 'rxjs/add/operator/do';
import { Observable } from 'rxjs/Observable';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';


export interface ConvertRequest {
  amount: string;
  from: string;
  to: string;
}

export interface ConvertResponse {
  result: string;
}

@Injectable()
export class ConverterService {
  constructor(
    private http: HttpClient,
  ) { }

  errorHandler(err: any) {
    console.log(`Converter service: error: ${err}`);
  }

  convert(convert: ConvertRequest): Observable<ConvertResponse> {
    return this.http
      .post<ConvertResponse>('/api/currencies/convert/', convert)
      .do(
        data => console.log(`convert success`),
        this.errorHandler,
      );
  }

}
