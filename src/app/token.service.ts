import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import * as types from './types';
import { SessionService } from './session.service';
import { environment } from '../environments/environment';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class TokenService {

  constructor(
    private ss: SessionService,
    private http: HttpClient
  ) { }

  transfer(from: string, to: string, amount: number): Promise<void> {
    return this.http
      .post<void>(environment.backendUrl + '/v1/token/transfer', {
        token: this.ss.getToken(),
        from: from,
        to: to,
        amount: amount
      }).toPromise();
  }

  createToken(tokenName: string, tokenDescription: string): Promise<void> {
    return this.http
      .post<void>(environment.backendUrl + '/v1/token', {
        token: this.ss.getToken(),
        serviceTokenName: tokenName,
        serviceTokenDescription: tokenDescription
      }).toPromise();
  }
}
