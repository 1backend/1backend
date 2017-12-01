import { Injectable } from '@angular/core';
import { Cookie } from 'ng2-cookies';

const d = new Date();
d.setFullYear(d.getFullYear() + 1);

@Injectable()
export class SessionService {
  constructor() {}
  setToken(t: string) {
    Cookie.set('token', t, d, '/');
  }
  getToken(): string {
    return Cookie.get('token');
  }
}
