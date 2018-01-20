import { Component, OnInit, Output, EventEmitter, Input } from '@angular/core';
import * as types from '../../../types';
import { FilterPipe } from '../../../filter.pipe';
import { Router } from '@angular/router';
import { UserService } from '../../../user.service';
import { ProjectService } from '../../../project.service';
import { environment } from '../../../../environments/environment';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-code',
  templateUrl: './code.component.html',
  styleUrls: ['./code.component.css']
})
export class CodeComponent implements OnInit {
  @Input() project: types.Project;
  @Output() onProjectSaved = new EventEmitter<void>();

  user: types.User;
  backendUrl = environment.backendUrl;
  search: string;
  endpoints: types.Endpoint[] = [];
  leftWidth = 49;
  rightWidth = 49;
  newEndpoint: types.Endpoint = {
    Method: 'GET',
    Url: '/',
    Description: ''
  };
  status: boolean;
  readme: string;
  currentPage = 0;
  readOnly = true;

  constructor(
    private router: Router,
    private us: UserService,
    private ps: ProjectService,
    private title: Title
  ) {
    this.user = this.us.user;
  }

  ngOnInit() {
    if (this.user.Nick === this.project.Author) {
      this.readOnly = false;
    }
  }

  add() {
    this.project.Endpoints.push(this.newEndpoint);
    this.save();
  }

  save() {
    this.ps
      .update(this.project)
      .then(data => {
        this.onProjectSaved.emit();
      })
      .catch(error => {});
  }

  getTestToken(): string {
    const tokens = this.us.user.Tokens;
    if (!tokens || !tokens.length) {
      return 'YOUR_API_KEY';
    }
    const testToken = tokens.filter(t => {
      return t.Name === 'test';
    });
    if (testToken.length === 0) {
      return 'YOUR_API_KEY';
    }
    return testToken[0].Token;
  }

  getAceCompatibleMode(): string {
    switch (this.project.Mode) {
      case 'go': {
        return 'golang';
      }
      case 'nodejs': {
        return 'javascript';
      }
      case 'typescript': {
        return 'typescript';
      }
    }
  }

  pageChanged($event: any) {
    this.currentPage = $event.pageIndex;
  }

  setRightFullScreen() {
    this.rightWidth = 80;
    this.leftWidth = 18;
  }
  smallScreen() {
    this.rightWidth = 49;
    this.leftWidth = 49;
  }
  setLeftFullScreen() {
    this.rightWidth = 19;
    this.leftWidth = 80;
  }

  isLeftFullScreen(): boolean {
    return this.leftWidth === 80;
  }

  isRightFullScreen(): boolean {
    return this.rightWidth === 80;
  }

  reveal(e: types.Endpoint) {
    e.Selected = !e.Selected;
  }

  goSql() {
    this.router.navigate([
      '/' + this.project.Author + '/' + this.project.Name + '/' + 'sql-console'
    ]);
    location.reload(); // XD
  }
}
