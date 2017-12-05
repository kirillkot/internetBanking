import { Component, OnInit, ChangeDetectorRef } from '@angular/core';

import { ManageListComponent } from '../abstract/manage-list.component';

import { CardForm, Card } from './card.service';
import { CardService } from './card.service';


@Component({
  selector: 'app-card-list',
  templateUrl: './card-list.component.html',
  styleUrls: ['./card-list.component.css']
})
export class CardListComponent extends ManageListComponent<CardForm, Card> {
  displayedColumns = ['name', 'ttl', 'type', 'status', 'currency', 'balance'];

  constructor(
    service: CardService,
    detector: ChangeDetectorRef,
  ) {
    super(service, detector);
  }

}
