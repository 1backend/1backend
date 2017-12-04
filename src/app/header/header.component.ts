import { Component, OnInit, Input, OnChanges } from '@angular/core';
import { UserService } from '../user.service';
import * as types from '../types';
import { LoginDialogService } from '../login/login-dialog.service';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  user: types.User;

  constructor(
    public us: UserService,
    private router: Router,
    private lds: LoginDialogService
  ) {
    this.user = this.us.user;
  }
  ngOnInit() {}

  login() {
    this.lds.openDialog(true, () => {
      this.router.navigate(['/' + this.us.user.Nick]);
    });
  }

  logout() {
    this.us.logout();
    this.router.navigate(['']);
    location.reload();
  }

  viewProfile() {
    this.router.navigate(['/' + this.user.Nick]);
  }

  getStarted() {
    this.router.navigate(['/', { fragment: 'getStarted' }]);
  }

  register() {
    this.lds.openDialog(false, () => {
      this.router.navigate(['/' + this.us.user.Nick]);
    });
  }

  pricing() {
    this.router.navigate(['/'], { fragment: 'pricing' });
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
