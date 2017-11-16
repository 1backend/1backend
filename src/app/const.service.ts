import { Injectable } from '@angular/core';

@Injectable()
export class ConstService {
  url: string;
  tokenMinLen: number;

  constructor() {
    this.url = window.location.hostname.includes('1backend')
      ? 'https://1backend.com:9993'
      : 'http://192.168.1.108:8883';
    this.tokenMinLen = 10;
  }
}
