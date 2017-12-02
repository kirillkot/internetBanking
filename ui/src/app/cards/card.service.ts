import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';


export interface Card {
  name: string;
  ttl: Date;
  type: string;
  status: string;
  currency: string;
  balance: string;
}

@Injectable()
export class CardService {

  constructor(
    private http: HttpClient,
  ) { }

  getCards(): Observable<User[]> {
    return this.http.get<User[]>('/api/cards');
  }

}
