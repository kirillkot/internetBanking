import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatStepper } from '@angular/material';

import { LoginService } from '../login.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginGroup: FormGroup;
  twofactorGroup: FormGroup;
  error: string;

  constructor(
    private location: Location,
    private service: LoginService,
    private formbuilder: FormBuilder,
  ) { }

  ngOnInit() {
    this.error = null;
    this.loginGroup = this.formbuilder.group({
      username: ['', [Validators.required, Validators.minLength(4)]],
      password: ['', [Validators.required, Validators.minLength(4)]],
    });
    this.twofactorGroup = this.formbuilder.group({
      username: [null, [Validators.required, Validators.minLength(4)]],
      password: [null, [Validators.required, Validators.minLength(4)]],
      twofactor: ['', [Validators.required]],
    });
  }

  login(steper: MatStepper) {
    return this.service.login(this.loginGroup.value)
      .subscribe(
        () => {
          this.error = null;
          this.twofactorGroup.patchValue(this.loginGroup.value);
          steper.next();
        },
        (err: any) => this.error = `Error: ${err}`,
      );
  }

  twofactor() {
    return this.service.twofactor(this.twofactorGroup.value)
      .subscribe(
        (data) => {
          this.error = null;
          this.location.back();
        },
        (err: any) => this.error = `Error: ${err}`,
      );
  }

}
