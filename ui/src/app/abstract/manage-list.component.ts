import { Observable } from 'rxjs/Observable';
import { Component, ChangeDetectorRef } from '@angular/core';

import { BackendService } from './backend.service';
import { BasicDataSource } from './data.source';


@Component({})
export abstract class ManageListComponent<Form, Model>{
  public dataSource: BasicDataSource<Model> = null;

  constructor(
    protected service: BackendService<Form, Model>,
    protected detector: ChangeDetectorRef,
  ) {
    this.dataSource = new BasicDataSource<Model>(service);
  }

  refresh(): void {
    this.dataSource = new BasicDataSource<Model>(this.service);
    this.detector.detectChanges();
  }

  delete(id: number): void {
    this.service.delete(id)
      .subscribe(
        data => this.refresh(),
      );
  }
}
