import { Component } from '@angular/core';


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
}
