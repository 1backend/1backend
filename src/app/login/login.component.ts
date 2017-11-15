import { Component, OnInit, Inject } from '@angular/core';
import { NotificationsService } from 'angular2-notifications';
import { HttpClient } from '@angular/common/http';
import { ConstService } from '../const.service';
import { SessionService } from '../session.service';
import { Router } from '@angular/router';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { UserService } from '../user.service';
import * as types from '../types';
import { FormControl, Validators } from '@angular/forms';

const EMAIL_REGEX = /^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/;

interface LoginResponse {
  token: types.AccessToken;
}

interface RegisterResponse {
  token: types.AccessToken;
}

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  username = '';
  email = '';
  password = '';
  password_conf = '';
  selectedIndex = 0;
  emailFormControl = new FormControl('', [Validators.pattern(EMAIL_REGEX)]);

  onNoClick(): void {
    this.dialogRef.close();
  }

  constructor(
    private notif: NotificationsService,
    private http: HttpClient,
    public dialogRef: MatDialogRef<LoginComponent>,
    private ss: SessionService,
    private _const: ConstService,
    private us: UserService,
    @Inject(MAT_DIALOG_DATA) public data: any,
    private router: Router
  ) {
    if (!this.data.isLogin) {
      this.selectedIndex = 1;
    }
  }


  login() {
    if (this.validator() !== 'ok') {
      return;
    }
    this.http
      .post<LoginResponse>(this._const.url + '/v1/login', {
        email: this.email,
        password: this.password
      })
      .subscribe(
        data => {
          this.ss.setToken(data.token.Token);
          // this.event.broadcast('logged-in')
          this.dialogRef.close();
          this.us.get().then(() => {
            if (this.data.callback()) {
            } else {
              this.router.navigate(['/' + this.us.user.Nick]);
            }
          });
        },
        error => {
          this.notif.error('Could not login');
        }
      );
  }

  register() {
    if (this.reg_validator() !== 'ok') {
      return;
    }
    this.http
      .post<RegisterResponse>(this._const.url + '/v1/register', {
        nick: this.username,
        email: this.email,
        password: this.password
      })
      .subscribe(
        data => {
          this.ss.setToken(data.token.Token);
          this.dialogRef.close();
          this.us.get().then(() => {
            this.data.callback();
          });
          setTimeout(() => {
            this.router.navigate(['/' + this.username]);
          }, 200); // database purposes
        },
        e => {
          this.notif.error(e.error.error);
        }
      );
  }

  validator() {
    if (!this.email) {
      this.notif.error('Email is empty');
      return;
    }
    if (!this.password) {
      this.notif.error('Password is incorrect');
      return;
    }
    return 'ok';
  }

  ngOnInit() {}

  reg_validator() {
    if (!this.username) {
      this.notif.error('Username is empty.');
      return;
    }
    if (!this.email.match('@')) {
      this.notif.error('Wrong email format');
      return;
    }
    if (!this.password) {
      this.notif.error('Password is empty.');
      return;
    }
    if (this.password_conf !== this.password) {
      this.notif.error('Passwords does not match');
      return;
    }
    return 'ok';
  }
}
