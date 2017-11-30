import { Component, OnInit, OnChanges, Input } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient, HttpParams } from '@angular/common/http';
import * as types from '../types';
import { Router, NavigationStart, Event } from '@angular/router';
import { SessionService } from '../session.service';
import { ChargeService } from '../charge.service';
import { ConstService } from '../const.service';
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
    let p = new HttpParams();
    p = p.set('nick', this.author);
    p = p.set('token', this.ss.getToken());
    this.http
      .get<types.Project[]>(this._const.url + '/v1/projects', {
        params: p
      })
      .subscribe(
        projs => {
          this.projects = projs.sort((a, b) => {
            if (a.UpdatedAt === b.UpdatedAt) {
              return 0;
            }
            if (a.UpdatedAt < b.UpdatedAt) {
              return 1;
            }
            return -1;
          });
        },
        error => {
          this.userFound = false;
        }
      );
    p = new HttpParams();
    p = p.set('token', this.ss.getToken());
    p = p.set('nick', this.author);
    this.http
      .get<types.User>(this._const.url + '/v1/user', { params: p })
      .subscribe(
        user => {
          this.user = user;
        },
        error => {
          this.userFound = false;
        }
      );
  }

  edit() {
    this.saved = false;
  }

  delete(project: types.Project) {
    this.http.delete(this._const.url + '/v1/project', {}).subscribe(
      data => {
        this.refresh();
      },
      error => {}
    );
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
