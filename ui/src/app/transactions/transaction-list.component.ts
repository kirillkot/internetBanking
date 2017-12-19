import { Location } from '@angular/common';
import { Component, ChangeDetectorRef, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';

import { SimpleDataSource } from '../abstract/data.source';

import { Transaction } from './transaction.service';
import { TransactionService } from './transaction.service';


@Component({
  selector: 'app-transaction-list',
  templateUrl: './transaction-list.component.html',
  styleUrls: ['./transaction-list.component.css']
})
export class TransactionListComponent implements OnInit {
  public displayedColumns = ['delta', 'time', 'detail'];
  public dataSource: SimpleDataSource<Transaction> = null;

  private account: number = null;

  constructor(
    private location: Location,
    private route: ActivatedRoute,
    private detector: ChangeDetectorRef,
    private service: TransactionService,
  ) { }

  refresh(): void {
    this.service.getAccountStat(this.account)
      .subscribe(
        (data) => {
          this.dataSource = new SimpleDataSource(data.transactions);
          this.detector.detectChanges();
        },
      );
  }

  ngOnInit() {
    this.account = +this.route.snapshot.paramMap.get('account');
    this.refresh();
  }

}
