import { Injectable } from '@angular/core';
import { environment } from '../environments/environment';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SessionService } from './session.service';
import { NotificationsService } from 'angular2-notifications';

const pricePer100k = 900;

@Injectable()
export class ChargeService {
  constructor(
    private http: HttpClient,
    private ss: SessionService,
    private ns: NotificationsService
  ) {}

  charge(amt: number, callback: () => void) {
    const that = this;
    const handler = (<any>window).StripeCheckout.configure({
      key: environment.stripeKey,
      locale: 'auto',
      token: function(token: any) {
        that.ns.alert('Payment in progress');
        that.http
          .post(environment.backendUrl + '/v1/charge', {
            paymentToken: token.id,
            token: that.ss.getToken(),
            amount: amt
          })
          .subscribe(
            data => {
              that.ns.success('Thank you', 'Your payment was successful');
              callback();
            },
            error => {
              that.ns.error(error);
            }
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
