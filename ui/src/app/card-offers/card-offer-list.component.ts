import { Component, OnInit, ChangeDetectorRef } from '@angular/core';

import { ManageListComponent } from '../abstract/manage-list.component';

import { CardOfferForm, CardOffer } from './card-offer.service';
import { CardOfferService } from './card-offer.service';


@Component({
  selector: 'app-card-offer-list',
  templateUrl: './card-offer-list.component.html',
  styleUrls: ['./card-offer-list.component.css']
})
export class CardOfferListComponent extends
    ManageListComponent<CardOfferForm, CardOffer> {
  displayedColumns = ['id', 'name', 'type', 'cashback', 'currency', 'ttlMonth', 'actions'];

  constructor(
    service: CardOfferService,
    detector: ChangeDetectorRef,
  ) {
    super(service, detector);
  }

}
