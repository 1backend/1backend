import { Injectable } from '@angular/core';

interface Callback {
  function: (data?: any) => void;
}

@Injectable()
export class EventService {
  subscribers: Map<string, Callback[]> = new Map<string, Callback[]>();
  x: number;
  constructor() {
    this.x = Math.random();
  }

  broadcast(channel: string, data?: any) {
    if (!this.subscribers.has(channel)) {
      return;
    }
    this.subscribers.get(channel).forEach(callback => {
      callback.function(data);
    });
  }

  // send this massage, but wait for the first subscriber to appear
  broadcastWhenListening(channel: string, data?: any, delay?: number) {
    if (!delay) {
      delay = 10;
    }
    setTimeout(() => {
      if (this.subscribers.has(channel)) {
        this.broadcast(channel, data);
      } else {
        this.broadcastWhenListening(channel, data, delay + 50);
      }
    }, delay);
  }

  subscribe(channel: string, f: (data: any) => void) {
    if (this.subscribers.has(channel)) {
      this.subscribers.get(channel).push({
        function: f
      });
    } else {
      this.subscribers = this.subscribers.set(channel, [{ function: f }]);
    }
  }
}
