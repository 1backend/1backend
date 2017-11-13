import { Component, OnInit, Input } from '@angular/core';
import { UserService } from '../user.service';
import * as types from '../types';
import { LoginDialogService } from '../login/login-dialog.service';
import { SessionService } from '../session.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  user: types.User;

  constructor(
    private us: UserService,
    private sess: SessionService,
    public ss: SessionService,
    private router: Router,
    private lds: LoginDialogService
  ) {
    this.user = this.us.user;
  }
  ngOnInit() {}

  login() {
    this.lds.openDialog(true, () => {});
  }
  logout() {
    this.ss.setToken('');
    this.router.navigate(['']);
    location.reload();
  }

  viewProfile() {
    this.router.navigate(['/' + this.user.Nick]);
  }

  register() {
    this.lds.openDialog(false, () => {});
  }

  scroll(el) {
    el.scrollIntoView();
  }

  discoverProjects() {
    this.router.navigate(['/projects']);
  }

  home() {
    this.router.navigate(['/']);
  }
}
