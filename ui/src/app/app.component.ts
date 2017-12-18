import { Router } from '@angular/router';
import { Component } from '@angular/core';

import { LoginService } from './login.service';


class RouterLink {
  constructor(
    public name: string,
    public link: string,
  ) {}
}

const AuthorizedLinks = [
    new RouterLink('Cards', '/cards'),
    new RouterLink('Payments', '/payments/management'),
]

const AdminLinks = AuthorizedLinks.concat([
    new RouterLink('Users', '/users'),
    new RouterLink('Accounts', '/accounts/management'),
    new RouterLink('Card Offers', '/card-offers/management'),
    new RouterLink('Payment Types', '/payment-types/management'),
    new RouterLink('Transactions', '/transactions/management'),
    new RouterLink('Currencies', '/currencies/management'),
])

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'app';

  constructor(
    private router: Router,
    private service: LoginService,
  ) { }

  getRoutes(): RouterLink[] {
    if ( this.service.isAdmin() ) {
      return AdminLinks;
    }
    return AuthorizedLinks;
  }

  logout(): void {
    this.service.logout();
    this.router.navigate(['/login']);
  }
}
