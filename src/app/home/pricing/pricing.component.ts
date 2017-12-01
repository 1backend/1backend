import { Component, OnInit } from '@angular/core';
import { LoginDialogService } from '../../login/login-dialog.service';
import { SessionService } from '../../session.service';
import { ChargeService } from '../../charge.service';
import { UserService } from '../../user.service';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Router } from '@angular/router';
import { environment } from '../../../environments/environment';
import * as types from '../../types';

const amt = 900;

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
    private charge: ChargeService
  ) {}

  ngOnInit() {}

  openCheckout() {
    const that = this;
    if (this.ss.getToken()) {
      this.charge.charge(amt, () => {
        that.router.navigate(['/' + that.us.user.Nick]);
      });
      return;
    }
    this.lds.openDialog(false, (tok: types.AccessToken) => {
      this.charge.charge(amt, () => {
        that.router.navigate(['/' + that.us.user.Nick]);
      });
    });
  }

  register() {
    this.lds.openDialog(false, () => {
      this.router.navigate(['/' + this.us.user.Nick]);
    });
  }
}
