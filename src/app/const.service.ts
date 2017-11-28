import { Injectable } from '@angular/core';
import { environment } from '../environments/environment';

@Injectable()
export class ConstService {
  url: string;
  tokenMinLen: number;

  constructor() {
    this.url = environment.backendUrl;
  }
}
