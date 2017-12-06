import { Router } from '@angular/router';
import { Component } from '@angular/core';

import { LoginService } from './login.service';


class RouterLink {
  constructor(
    public name: string,
    public link: string,
  ) {}
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'app';
  routes = [
    new RouterLink('Users', '/users'),
    new RouterLink('Login', '/login'),
    new RouterLink('Cards', '/cards'),
    new RouterLink('Card Offers', '/card-offers/management'),
  ];

  constructor(
    private router: Router,
    private service: LoginService,
  ) { }

  logout(): void {
    this.service.logout();
    this.router.navigate(['/login']);
  }
}
