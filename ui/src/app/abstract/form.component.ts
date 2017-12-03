import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';

import { BackendService } from './backend.service';


@Component({})
export abstract class FormComponent<Form, Model> implements OnInit {
  public group: FormGroup;

  constructor(
    protected location: Location,
    protected formbuilder: FormBuilder,
    protected service: BackendService<Form, Model>,
  ) { }

  abstract fields(): any;

  ngOnInit() {
    this.group = this.formbuilder.group(this.fields());
  }

  create(): void {
    this.service.create(this.group.value);
    this.back();
  }

  back(): void {
    this.location.back();
  }

}
