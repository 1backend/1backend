import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-pricing',
  templateUrl: './pricing.component.html',
  styleUrls: ['./pricing.component.css']
})
export class PricingComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

  openCheckout(amount: number) {
    const handler = (<any>window).StripeCheckout.configure({
      key: 'pk_test_jN1awbuFc99uOJvciajTVvCU',
      locale: 'auto',
      token: function (token: any) {
        // You can access the token ID with `token.id`.
        // Get the token ID to your server-side code for use.
      }
    });

    handler.open({
      name: 'Starter',
      description: '',
      amount: 19000
    });

  }
}
