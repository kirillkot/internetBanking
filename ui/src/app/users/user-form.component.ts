import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

import {
  User,
  UserForm,
  UserService
} from './user.service';


@Component({
  selector: 'app-user-form',
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.css']
})
export class UserFormComponent implements OnInit {
  user: FormGroup;

  constructor(
    private location: Location,
    private service: UserService,
    private formbuilder: FormBuilder,
  ) { }

  ngOnInit() {
    this.user = this.formbuilder.group({
      username: ['', [Validators.required, Validators.minLength(4)]],
      isAdmin: [false, [Validators.required]],
      password: ['', [Validators.required, Validators.minLength(4)]],
    });
    console.log(`User Form: init user: ${this.user}`)
  }

  create(): void {
    console.log(`User Form: onSubmit: ${this.user}`)
    let newuser = this.service.create(this.user.value);
  }

  back(): void {
    this.location.back();
  }

}
