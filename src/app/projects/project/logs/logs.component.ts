import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import * as types from '../../../types';
import { ProjectService } from '../../../project.service';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from '../../../../environments/environment';
import { SessionService } from '../../../session.service';

interface ReloadResponse {
  Logs: string;
}

@Component({
  selector: 'app-logs',
  templateUrl: './logs.component.html',
  styleUrls: ['./logs.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class LogsComponent implements OnInit {
  @Input() project: types.Project;
  logs: string;
  constructor(
    private ps: ProjectService,
    private http: HttpClient,
    private ss: SessionService
  ) {}

  ngOnInit() {
    this.reload();
  }

  reload() {
    let p = new HttpParams();
    p = p.set('projectId', this.project.Id);
    p = p.set('token', this.ss.getToken());
    this.http
      .get<ReloadResponse>(environment.backendUrl + '/v1/logs', {
        params: p
      })
      .subscribe(
        data => {
          this.logs = data.Logs;
        },
        error => {}
      );
  }
}
