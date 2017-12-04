import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';
import * as types from '../../types';
import { UserService } from '../../user.service';
import { ProjectService } from '../../project.service';
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
    public us: UserService,
    private ps: ProjectService,
    private router: Router,
    private loginDialog: LoginDialogService,
    private notif: NotificationsService
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

    this.ps.save(p).then(() => {
      if (this.callback) {
        this.callback(p);
      } else {
        this.router.navigate(['/' + this.user.Nick + '/' + this.name]);
      }
    });
  }

  create() {
    if (!this.us.loggedIn) {
      this.loginDialog.openDialog(true, () => {
        this.createProject();
      });
    } else {
      this.createProject();
    }
  }
}
