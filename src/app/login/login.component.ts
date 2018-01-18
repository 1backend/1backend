import { Component, OnInit, Inject } from '@angular/core';
import { NotificationsService } from 'angular2-notifications';
import { Router } from '@angular/router';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { UserService } from '../user.service';
import * as types from '../types';
import { FormControl, Validators } from '@angular/forms';

const EMAIL_REGEX = /^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,63}$/;
const USERNAME_REGEX = /^[a-z0-9]+([a-z0-9\-]*)*[a-z0-9]+$/;

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
  usernameFormControl = new FormControl('', [
    Validators.pattern(USERNAME_REGEX)
  ]);

  onNoClick(): void {
    this.dialogRef.close();
  }

  constructor(
    private notif: NotificationsService,
    public dialogRef: MatDialogRef<LoginComponent>,
    private us: UserService,
    @Inject(MAT_DIALOG_DATA) public data: any,
    private router: Router
  ) {
    if (!this.data.isLogin) {
      this.selectedIndex = 1;
    }
  }

  login() {
    if (!this.loginValid()) {
      return;
    }
    this.us
      .login(this.email, this.password)
      .then(loginResponse => {
        // this.event.broadcast('logged-in')
        this.dialogRef.close();
        this.us.get().then(() => {
          if (this.data.callback) {
            this.data.callback(loginResponse.token);
          } else {
            this.router.navigate(['/' + this.us.user.Nick]);
          }
        });
      })
      .catch(error => {
        this.notif.error('Could not login');
      });
  }

  register() {
    if (!this.registerValid()) {
      return;
    }
    this.us.register(this.username, this.email, this.password).then(
      registerResponse => {
        this.dialogRef.close();
        this.us.get().then(() => {
          if (this.data.callback) {
            this.data.callback(registerResponse.token);
          } else {
            this.router.navigate(['/' + this.us.user.Nick]);
          }
        });
      },
      e => {
        this.notif.error(e.error.error);
      }
    );
  }

  loginValid(): boolean {
    if (!this.email) {
      this.notif.error('Email is empty');
      return false;
    }
    if (!this.password) {
      this.notif.error('Password is incorrect');
      return false;
    }
    return true;
  }

  ngOnInit() {}

  registerValid(): boolean {
    if (!this.username) {
      this.notif.error('Username is empty.');
      return false;
    }
    if (this.username === 'config') {
      this.notif.error('"config" is not allowed as username.');
      return false;
    }
    if (this.username === 'projects') {
      this.notif.error('"projects" is not allowed as username.');
      return false;
    }
    if (this.username === 'reset') {
      this.notif.error('"reset" is not allowed as username.');
      return false;
    }
    if (this.username === 'recover') {
      this.notif.error('"recover" is not allowed as username.');
      return false;
    }
    if (!USERNAME_REGEX.test(this.username)) {
      this.notif.error(
        'Username must contain only lowercase and numerical characters.'
      );
      return false;
    }
    if (!EMAIL_REGEX.test(this.email)) {
      this.notif.error('Email is invalid.');
      return false;
    }
    if (!this.password) {
      this.notif.error('Password is empty.');
      return false;
    }
    if (this.password_conf !== this.password) {
      this.notif.error('Passwords does not match.');
      return false;
    }
    return true;
  }
}
