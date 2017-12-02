import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

import { UserService } from './user.service';


@Component({
  selector: 'app-user-login',
  templateUrl: './user-login.component.html',
  styleUrls: ['./user-login.component.css']
})
export class UserLoginComponent implements OnInit {
  user: FormGroup;

  constructor(
    private location: Location,
    private service: UserService,
    private formbuilder: FormBuilder,
  ) { }

  ngOnInit() {
    this.user = this.formbuilder.group({
      username: ['', [Validators.required, Validators.minLength(4)]],
      password: ['', [Validators.required, Validators.minLength(4)]],
    });
    console.log(`User Login Form: init user: ${this.user}`)
  }

  login(): void {
    console.log(`User Login Form: create: ${this.user}`);
    this.service.login(this.user.value);
    // this.location.back();
  }

}
