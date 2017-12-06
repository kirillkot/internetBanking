import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

import { LoginService } from '../login.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  group: FormGroup;
  error: string;

  constructor(
    private location: Location,
    private service: LoginService,
    private formbuilder: FormBuilder,
  ) { }

  ngOnInit() {
    this.error = null;
    this.group = this.formbuilder.group({
      username: ['', [Validators.required, Validators.minLength(4)]],
      password: ['', [Validators.required, Validators.minLength(4)]],
    });
  }

  login(): void {
    this.service.login(this.group.value)
      .subscribe(
        () => {
          this.error = null;
          this.location.back();
        },
        (err: any) => this.error = `Error: ${err}`,
      );
  }

}
