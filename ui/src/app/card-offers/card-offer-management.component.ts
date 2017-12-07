import { Component, ChangeDetectorRef } from '@angular/core';

import { ManageListComponent } from '../abstract/manage-list.component';

import { CardOfferForm, CardOffer } from './card-offer.service';
import { CardOfferService } from './card-offer.service';


@Component({
  selector: 'app-card-offer-management',
  templateUrl: './card-offer-management.component.html',
  styleUrls: ['./card-offer-management.component.css']
})
export class CardOfferManagementComponent extends
    ManageListComponent<CardOfferForm, CardOffer> {
  displayedColumns = ['id', 'name', 'type', 'cashback', 'currency', 'ttlMonth', 'actions'];

  constructor(
    service: CardOfferService,
    detector: ChangeDetectorRef,
  ) {
    super(service, detector);
  }

}
