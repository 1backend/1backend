import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { SessionService } from '../session.service';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { ActivatedRoute, Params, Router } from '@angular/router';
import * as types from '../types';
import { NotificationsService } from 'angular2-notifications';
import { Title } from '@angular/platform-browser';

interface ResetPasswordResponse {
  token: types.AccessToken;
}

@Component({
  selector: 'app-reset',
  templateUrl: './reset.component.html',
  styleUrls: ['../login/login.component.css']
})
export class ResetComponent implements OnInit {
  newPassword: string;
  newPassword2: string;
  secret: string;

  constructor(
    private ss: SessionService,
    private http: HttpClient,
    private activatedRoute: ActivatedRoute,
    private router: Router,
    private notif: NotificationsService,
    private title: Title
  ) {}

  ngOnInit() {
    this.title.setTitle('Password reset');
    this.secret = this.activatedRoute.snapshot.params['secret'];
  }

  resetPassword() {
    if (this.newPassword !== this.newPassword2) {
      this.notif.error('The two passwords do not match');
      return;
    }
    this.http
      .post<ResetPasswordResponse>(environment.backendUrl + '/v1/reset-password', {
        newPassword: this.newPassword,
        secret: this.secret
      })
      .subscribe(
        rsp => {
          if (!rsp.token.Token) {
            this.notif.error('Can\'t find token');
          } else {
            this.ss.setToken(rsp.token.Token);
            this.router.navigate(['']);
          }
        },
        err => {
          this.notif.error('Something went wrong');
        }
      );
  }
}
