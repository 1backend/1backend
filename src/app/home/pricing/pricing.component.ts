import { Component, OnInit } from '@angular/core';
import { LoginDialogService } from '../../login/login-dialog.service';
import { environment } from '../../../environments/environment';
import * as types from '../../types';

@Component({
  selector: 'app-pricing',
  templateUrl: './pricing.component.html',
  styleUrls: ['./pricing.component.css']
})
export class PricingComponent implements OnInit {
  constructor(private lds: LoginDialogService) {}

  ngOnInit() {}

  openCheckout(amount: number) {
    const amt = 1900;
    this.lds.openDialog(false, (tok: types.AccessToken) => {
      const handler = (<any>window).StripeCheckout.configure({
        key: environment.stripeKey,
        locale: 'auto',
        token: function(token: any) {
          this.http
            .poste(this._const.url + '/v1/charge', {
              paymentToken: token.id,
              token: tok.Token,
              amount: amt
            })
            .subscribe(
              data => {
                this.refresh();
              },
              error => {}
            );
        }
      });
      handler.open({
        name: 'Starter',
        description: '',
        amount: amt
      });
    });
  }

  register() {
    this.lds.openDialog(false, () => {});
  }
}
