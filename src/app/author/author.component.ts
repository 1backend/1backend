import { Component, OnInit, OnChanges, Input } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import * as types from '../types';
import { Router, NavigationStart, Event } from '@angular/router';
import { ChargeService } from '../charge.service';
import { ProjectService } from '../project.service';
import { NotificationsService } from 'angular2-notifications';
import { UserService } from '../user.service';
import { EventService } from '../event.service';

@Component({
  selector: 'app-author',
  templateUrl: './author.component.html',
  styleUrls: ['./author.component.css']
})
export class AuthorComponent implements OnInit {
  amount = 9;
  author = '';
  name = '';
  projects: types.Project[] = [];
  saved = true;
  user: types.User = {} as types.User;
  userFound = true;
  selectedTabIndex = 0; // workaround for bug https://github.com/angular/material2/issues/5269

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private notif: NotificationsService,
    public us: UserService,
    public ps: ProjectService,
    private charge: ChargeService
  ) {
    this.refresh();
  }

  save() {
    if (!this.valid()) {
      return;
    }
    this.us.saveSelf().subscribe(
      data => {
        this.saved = true;
        this.refresh();
      },
      error => {}
    );
  }

  ngOnInit() {
    // this solution is a bit crufty, but works fine.
    this.router.events.subscribe((event: Event) => {
      if (event instanceof NavigationStart) {
        setTimeout(() => {
          this.refresh();
        }, 500);
      }
    });
  }

  refresh() {
    this.author = this.route.snapshot.params['author'];
    this.ps
      .listByNick(this.author)
      .then(projects => (this.projects = projects))
      .catch(err => (this.userFound = false));
    this.us
      .getByNick(this.author)
      .then(user => (this.user = user))
      .catch(err => (this.userFound = false));
  }

  edit() {
    this.saved = false;
  }

  delete(project: types.Project) {
  }

  getAllCallsLeft(): number {
    if (!this.us.user.Tokens) {
      return 0;
    }
    return this.us.user.Tokens.map(t => t.Quota).reduce((prev, curr) => {
      return prev + curr;
    });
  }

  valid(): boolean {
    if (!this.user.Nick) {
      this.notif.error('Nickname is empty');
      return true;
    }
    if (!this.user.Email) {
      this.notif.error('Email is empty');
      return true;
    }
    return false;
  }

  purchase() {
    const that = this;
    this.charge.charge(this.amount * 100, () => {
      that.us.get();
    });
  }
}
