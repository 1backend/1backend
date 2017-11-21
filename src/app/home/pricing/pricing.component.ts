import { Component, OnInit } from '@angular/core';
import { LoginDialogService } from '../../login/login-dialog.service';
import { SessionService } from '../../session.service';
import { ConstService } from '../../const.service';
import { UserService } from '../../user.service';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Router } from '@angular/router';
import { environment } from '../../../environments/environment';
import * as types from '../../types';

const amt = 1900;

@Component({
  selector: 'app-pricing',
  templateUrl: './pricing.component.html',
  styleUrls: ['./pricing.component.css']
})
export class PricingComponent implements OnInit {
  constructor(
    private lds: LoginDialogService,
    public ss: SessionService,
    private http: HttpClient,
    private router: Router,
    private us: UserService,
    private _const: ConstService
  ) {}

  ngOnInit() {}

  openCheckout() {
    if (this.ss.getToken()) {
      this.charge(this.ss.getToken());
      return;
    }
    this.lds.openDialog(false, (tok: types.AccessToken) => {
      this.charge(tok.Token);
    });
  }

  charge(tok: string) {
    const that = this;
    const handler = (<any>window).StripeCheckout.configure({
      key: environment.stripeKey,
      locale: 'auto',
      token: function(token: any) {
        that.http
          .post(that._const.url + '/v1/charge', {
            paymentToken: token.id,
            token: tok,
            amount: amt
          })
          .subscribe(
            data => {
              that.router.navigate(['/' + that.us.user.Nick]);
            },
            error => {}
          );
      }
    });
    handler.open({
      name: 'Buy 100k quota',
      description: '',
      amount: amt
    });
  }

  register() {
    this.lds.openDialog(false, () => {
      this.router.navigate(['/' + this.us.user.Nick]);
    });
  }
}
