import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import * as types from './types';
import { SessionService } from './session.service';
import { ConstService } from './const.service';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class TokenService {

  constructor(
    private ss: SessionService,
    private _const: ConstService,
    private http: HttpClient
  ) { }

  transfer(from: string, to: string, amount: number): Observable<Object> {
    return this.http
      .post(this._const.url + '/v1/token/transfer', {
        token: this.ss.getToken(),
        from: from,
        to: to,
        amount: amount
      });
  }

  createToken(tokenName: string, tokenDescription: string): Observable<void> {
    return this.http
      .post<void>(this._const.url + '/v1/token', {
        token: this.ss.getToken(),
        serviceTokenName: tokenName,
        serviceTokenDescription: tokenDescription
      });
  }
}
