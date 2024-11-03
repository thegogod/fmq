import { Injectable } from '@angular/core';
import { TopicState } from './topics';

@Injectable({ providedIn: 'root' })
export class State {
  get json() {
    return {
      topics: this.topics$.json
    };
  }

  readonly topics$ = new TopicState();
}
