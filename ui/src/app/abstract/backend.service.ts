import 'rxjs/add/operator/do';
import { Observable } from 'rxjs/Observable';


import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpErrorResponse } from '@angular/common/http';

import { IObjectsService } from './data.source';


@Injectable()
export class BackendService<Form, Model> implements IObjectsService<Model> {
  constructor(
    protected entry: string,
    protected http: HttpClient,
  ) { }

  errorHandler(err: any) {
    console.log(`${this.entry} service: error: ${err}`);
  }

  create(object: Form): Observable<Model> {
    return this.http
      .post<Model>(`/api/${this.entry}/`, object)
      .do(
        data => console.log(`create ${this.entry}: success ${data}`),
        this.errorHandler,
      );
  }

  get(id: number): Observable<Model> {
    return this.http.get<Model>(`/api/${this.entry}/${id}`);
  }

  getObjects(): Observable< Array<Model> > {
    return this.http.get< Array<Model> >(`/api/${this.entry}/`);
  }

  delete(id: number): Observable<Model> {
    console.log(`delete ${this.entry}: with ${id}`);
    return this.http
      .delete<Model>(`/api/${this.entry}/${id}/`)
      .do(
        data => console.log(`Delete user: success ${id}`),
        this.errorHandler,
      );
  }

}
