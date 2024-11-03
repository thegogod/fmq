import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';

import { Topic } from '../models';

export class Topics {
  constructor(private readonly _http: HttpClient) { }

  async get() {
    return firstValueFrom(this._http.get<Record<string, Topic>>('/v1/topics'));
  }

  async getOne(name: string) {
    return firstValueFrom(this._http.get<Topic>(`/v1/topics/${name}`))
  }
}
