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
  displayedColumns = ['name', 'ttl', 'type', 'status',
      'currency', 'balance', 'actions'];

  constructor(
    detector: ChangeDetectorRef,
    private cardservice: CardService,
  ) {
    super(cardservice, detector);
  }

  nextState(card: Card): string {
    if ( card.status === "active" ) {
      return "Block";
    }
    return "Activate";
  }

  block(card: Card) {
    return this.cardservice.block(card.id)
      .subscribe(
        (data) => this.refresh(),
      );
  }
}
