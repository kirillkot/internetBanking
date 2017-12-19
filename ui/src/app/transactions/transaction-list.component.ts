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
  public dataSource: SimpleDataSource<Transaction> = null;
  public displayedColumns = ['delta', 'time', 'detail'];

  private card: number = null;

  constructor(
    private location: Location,
    private route: ActivatedRoute,
    private detector: ChangeDetectorRef,
    private service: TransactionService,
  ) { }

  refresh(): void {
    this.service.getObjectsByCard(this.card)
      .subscribe(
        (data) => {
          this.dataSource = new SimpleDataSource(data);
          this.detector.detectChanges();
        },
      );
  }

  ngOnInit() {
    this.card = +this.route.snapshot.paramMap.get('card');
    this.refresh();
  }

}
