import { Injectable } from '@angular/core';
import { environment } from '../environments/environment';
import { HttpClient, HttpParams } from '@angular/common/http';
import { ConstService } from './const.service';
import { SessionService } from './session.service';

const pricePer100k = 1900;

@Injectable()
export class ChargeService {
  constructor(
    private http: HttpClient,
    private _const: ConstService,
    private ss: SessionService
  ) {}

  charge(amt: number, callback: () => void) {
    const that = this;
    const handler = (<any>window).StripeCheckout.configure({
      key: environment.stripeKey,
      locale: 'auto',
      token: function(token: any) {
        that.http
          .post(that._const.url + '/v1/charge', {
            paymentToken: token.id,
            token: that.ss.getToken(),
            amount: amt
          })
          .subscribe(
            data => {
              callback();
            },
            error => {}
          );
      }
    });
    handler.open({
      name: 'Buy ' + Math.floor(100000 * amt / pricePer100k) + ' quota',
      description: '',
      amount: amt
    });
  }
}
