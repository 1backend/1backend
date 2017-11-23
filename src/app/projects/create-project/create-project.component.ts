import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';
import * as types from '../../types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { UserService } from '../../user.service';
import { SessionService } from '../../session.service';
import { ConstService } from '../../const.service';
import { LoginDialogService } from '../../login/login-dialog.service';
import { NotificationsService } from 'angular2-notifications';
import { FormControl, Validators } from '@angular/forms';

const pNAME_REGEX = /^[a-z]+([-]?[a-z0-9]+[-]?)+[a-z0-9]+$/;

@Component({
  selector: 'app-create-project',
  templateUrl: './create-project.component.html',
  styleUrls: ['./create-project.component.css']
})
export class CreateProjectComponent implements OnInit {
  @Input() callback: (p: types.Project) => void;
  name = '';
  user: types.User;
  private = false;

  languages: types.Languages[] = [
    {
      Name: 'nodejs',
      DisplayName: 'NodeJS',
      DisplayColor: '#80bd01'
    },
    {
      Name: 'go',
      DisplayName: 'Go',
      DisplayColor: '#4dd0e1'
    },
    {
      Name: 'typescript',
      DisplayName: 'TypeScript',
      DisplayColor: '#0074c1'
    }
  ];
  dependencies: types.Dependency[] = [
    {
      Type: 'mysql',
      DisplayName: 'MySQL'
    }
  ];

  constructor(
    private http: HttpClient,
    public us: UserService,
    private router: Router,
    private ss: SessionService,
    private loginDialog: LoginDialogService,
    private notif: NotificationsService,
    private _const: ConstService
  ) {
    this.user = this.us.user;
  }

  projectNameControl = new FormControl('', [Validators.pattern(pNAME_REGEX)]);

  resetLanguages() {
    this.languages.forEach(l => {
      l.Selected = false;
    });
  }

  ngOnInit() {}

  createProject() {
    if (!pNAME_REGEX.test(this.name) || this.name.length < 1) {
      return;
    }
    const lang = this.languages.filter(l => {
      return l.Selected;
    })[0];
    if (!lang) {
      this.notif.error('Please select a project type');
      return;
    }
    const dep = this.dependencies.filter(l => {
      return l.Selected;
    });
    if (!dep || dep.length < 1) {
      this.notif.error('Please select dependencies');
      return;
    }

    const p = {
      Author: this.user.Nick,
      Name: this.name,
      Mode: lang.Name,
      Dependencies: dep,
      Public: !this.private,
      OpenSource: true
    };

    this.http
      .post(this._const.url + '/v1/project', {
        token: this.ss.getToken(),
        project: p,
      })
      .subscribe(
        data => {
          if (this.callback) {
            this.callback(p);
          } else {
            this.router.navigate(['/' + this.user.Nick + '/' + this.name]);
          }
        },
        error => {
          this.notif.error(error.error.error);
          // alert(JSON.stringify(error));
        }
      );
  }

  create() {
    if (this.ss.getToken().length === 0) {
      this.loginDialog.openDialog(true, () => {
        this.createProject();
      });
    } else {
      this.createProject();
    }
  }
}
