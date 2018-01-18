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
    ApiGeneration: {},
    NpmPublication: {},
    Sitemap: {}
  };

  constructor(private http: HttpClient, private ss: SessionService) {}

  ngOnInit() {
    this.getConfig();
  }

  getConfig() {
    let p = new HttpParams();
    p = p.set('token', this.ss.getToken());
    this.http
      .get<types.Config>(environment.backendUrl + '/v1/config', {
        params: p
      })
      .toPromise()
      .then(conf => {
        this.config = conf;
      })
      .catch(err => (err = console.log('error')));
  }

  saveConfig() {
    this.http
      .put(environment.backendUrl + '/v1/config', {
        config: this.config,
        token: this.ss.getToken()
      })
      .toPromise()
      .then(() => this.getConfig())
      .catch(err => (err = console.log('error')));
    console.log(this.config.Sitemap.Enabled);
    console.log(this.config.Sitemap.Path);
  }
}
