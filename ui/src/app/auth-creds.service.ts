import { Injectable } from '@angular/core';

export interface AuthCreds {
  username: string;
  password: string;
}

@Injectable()
export class AuthCredsService {
  creds: AuthCreds = null;
  admin: boolean;

  setCreds(creds: AuthCreds | null, admin: boolean): void {
    this.creds = creds;
    this.admin = admin;
  }

  buildAuthHeader(creds: AuthCreds): string {
    return `Basic ${btoa(creds.username + ':' + creds.password)}`;
  }

  isAdmin(): boolean {
    if (this.creds === null || this.admin === null) {
      return false;
    }
    return this.admin;
  }

  getAuthHeader(): string {
    if (this.creds === null) {
      return null;
    }
    return this.buildAuthHeader(this.creds);
  }

}
