import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

import { AccountService } from './account.service';


@Component({
  selector: 'app-account-details',
  templateUrl: './account-details.component.html',
  styleUrls: ['./account-details.component.css']
})
export class AccountDetailsComponent implements OnInit {
  public group: FormGroup;

  constructor(
    private location: Location,
    private formbuilder: FormBuilder,
    private route: ActivatedRoute,
    private service: AccountService,
  ) { }

  ngOnInit() {
    const id = this.route.snapshot.paramMap.get('id');
    this.group = this.formbuilder.group({
      id: [id, [Validators.required]],
      amount: [0, [Validators.required]],
    });
  }

  addFunds(): void {
    this.service.addFunds(this.group.value)
      .subscribe(
        (data) => this.back(),
      );
  }

  back() {
    this.location.back();
  }
}
