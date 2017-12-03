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

  create(object: Form): Model {
    let result: Model = null;
    this.http
      .post<Model>(`/api/${this.entry}/`, object)
      .subscribe(
        data => {
          console.log(`create ${this.entry}: success ${data}`);
          result = data;
        },
        this.errorHandler,
      );
    return result;
  }

  get(id: number): Observable<Model> {
    return this.http.get<Model>(`/api/${this.entry}/${id}`);
  }

  getObjects(): Observable< Array<Model> > {
    return this.http.get< Array<Model> >(`/api/${this.entry}/`);
  }

  delete(id: number): void {
    console.log(`delete ${this.entry}: with ${id}`);
    this.http
      .delete<Model>(`/api/${this.entry}/${id}/`)
      .subscribe(
        data => {
          console.log(`Delete user: success ${id}`)
        },
        this.errorHandler,
      );
  }

}
