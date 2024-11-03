import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Topics } from './topics';

@Injectable({ providedIn: 'root' })
export class Api {
  readonly topics: Topics;

  constructor(private readonly _http: HttpClient) {
    this.topics = new Topics(this._http);
  }
}
