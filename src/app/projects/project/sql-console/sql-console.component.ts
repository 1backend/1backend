import { Component, OnInit, Input } from '@angular/core';
import * as types from '../../../types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SessionService } from '../../../session.service';
import { ConstService } from '../../../const.service';

interface RunSqlResponse {
  Answer: string;
}

@Component({
  selector: 'app-sql-console',
  templateUrl: './sql-console.component.html',
  styleUrls: ['./sql-console.component.css']
})
export class SqlConsoleComponent implements OnInit {
  @Input() project: types.Project;
  sqlConsole = 'SELECT CURRENT_USER();';
  sqlAnswer = '';

  constructor(
    private http: HttpClient,
    private ss: SessionService,
    private _const: ConstService
  ) {}

  ngOnInit() {}

  runSql() {
    this.http
      .post<RunSqlResponse>(this._const.url + '/v1/run-sql', {
        projectId: this.project.Id,
        sql: this.sqlConsole,
        token: this.ss.getToken()
      })
      .subscribe(
        data => {
          this.sqlAnswer = data.Answer;
        },
        error => {}
      );
  }
}
