import { Component, OnInit, Input } from '@angular/core';
import * as types from '../../../types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SessionService } from '../../../session.service';
import { environment } from '../../../../environments/environment';

interface RunSqlResponse {
  Answer: string;
}

@Component({
  selector: 'app-sql',
  templateUrl: './sql.component.html',
  styleUrls: ['./sql.component.css']
})
export class SqlComponent implements OnInit {
  @Input() project: types.Project;
  sqlConsole = 'SELECT CURRENT_USER();';
  sqlAnswer = '';

  constructor(
    private http: HttpClient,
    private ss: SessionService
  ) {}

  ngOnInit() {}

  runSql() {
    this.http
      .post<RunSqlResponse>(environment.backendUrl + '/v1/run-sql', {
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
