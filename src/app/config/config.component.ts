import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import * as types from '../types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { MatInputModule } from '@angular/material';
import { SessionService } from '../session.service';
@Component({
  selector: 'app-config',
  templateUrl: './config.component.html',
  styleUrls: ['./config.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class ConfigComponent implements OnInit {
  config: types.Config = {
    Api: {},
    Npm: {},
    Site: {}
  };

  constructor(private http: HttpClient, private ss: SessionService) {}

  ngOnInit() {}

  saveConfig() {
    this.http
      .post(environment.backendUrl + '/v1/config', {
        config: this.config,
        token: this.ss.getToken()
      })
      .toPromise()
      .catch(err => (err = console.log('error')));
  }
}
