import 'rxjs/add/observable/of';
import { Observable } from 'rxjs/Observable';
import { DataSource } from '@angular/cdk/collections';


export interface IObjectsService<Model> {
  getObjects(): Observable< Array<Model> >;
}

export class BasicDataSource<Model> extends DataSource<any> {
  constructor(
    protected service: IObjectsService<Model>
  ) {
    super();
  }

  connect(): Observable< Array<Model> > {
    return this.service.getObjects();
  }

  disconnect() {}
}

export class SimpleDataSource<Model> extends DataSource<any> {
  constructor(
    private objects: Array<Model>,
  ) {
    super();
  }

  connect(): Observable< Array<Model> > {
    return Observable.of(this.objects);
  }

  disconnect() {}

}
