import { Location } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { FormComponent } from '../abstract/form.component';

import { UserForm, User, UserService } from './user.service';


@Component({
  selector: 'app-user-form',
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.css']
})
export class UserFormComponent extends FormComponent<UserForm, User>{
  constructor(
    location: Location,
    formbuilder: FormBuilder,
    service: UserService,
  ) {
    super(location, formbuilder, service);
  }

  fields(): any {
    return {
      username: ['', [Validators.required, Validators.minLength(4)]],
      first_name: ['', [Validators.required, Validators.minLength(2)]],
      last_name: ['', [Validators.required, Validators.minLength(2)]],
      city_name: ['', [Validators.required, Validators.minLength(2)]],
      adress: ['', [Validators.required, Validators.minLength(2)]],
      isAdmin: [false, [Validators.required]],
      password: ['', [Validators.required, Validators.minLength(4)]],
    };
  }

}
