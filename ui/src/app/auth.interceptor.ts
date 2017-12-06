import 'rxjs/add/operator/do';

import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';
import { Injectable } from '@angular/core';
import { HttpEvent, HttpErrorResponse } from '@angular/common/http';
import { HttpInterceptor, HttpHandler, HttpRequest } from '@angular/common/http';

import { AuthCredsService } from './auth-creds.service';


@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  constructor(
    private service: AuthCredsService,
    private router: Router,
  ) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    // Get the auth header from the service.
    const authHeader = this.service.getAuthHeader()
    if (authHeader !== null && !req.headers.has('Authorization') ) {
      // Clone the request to add the new header.
      req = req.clone({headers: req.headers.set('Authorization', authHeader)});
    }

    // Pass on the cloned request instead of the original request.
    return next.handle(req).do(
      (event: HttpEvent<any>) => {},
      (err: any) => {
        if (err instanceof HttpErrorResponse) {
          if (err.status === 401 || err.status === 403) {
            this.router.navigate(['/login']);
          }
        }
      }
    );
  }
}
