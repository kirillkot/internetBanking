import { Observable } from 'rxjs/Observable';
import { DataSource } from '@angular/cdk/collections';
import { Component, OnInit } from '@angular/core';

import { Card } from './card.service';
import { CardService } from './card.service';


@Component({
  selector: 'app-card-list',
  templateUrl: './card-list.component.html',
  styleUrls: ['./card-list.component.css']
})
export class CardListComponent implements OnInit {
  displayedColumns = ['name', 'ttl', 'type', 'status', 'currency', 'balance'];
  dataSource: CardsDataSource;

  constructor(
    private service: CardService,
  ) {
    this.dataSource = new CardsDataSource(service);
  }

  ngOnInit() {
  }

}

export class CardsDataSource extends DataSource<any> {
  constructor(private service: CardService) {
    super();
  }
  connect(): Observable<[]> {
    return this.service.getCards();
  }
  disconnect() {}
}
