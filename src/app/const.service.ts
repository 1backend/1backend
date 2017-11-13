import { Injectable } from '@angular/core';

@Injectable()
export class ConstService {
  url: string;
  tokenMinLen: number;

  constructor() {
    this.url = window.location.hostname.includes('1backend')
      ? 'http://api.1backend.com'
      : 'http://192.168.0.10:8883';
    this.tokenMinLen = 10;
  }
}
