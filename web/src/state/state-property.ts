import { BehaviorSubject } from 'rxjs';

export class StateProperty<T = any> extends BehaviorSubject<T> {
  get json() {
    const value: any = { };

    for (const key in this.value) {
      value[key] = this.value[key];

      if (value[key] instanceof StateProperty) {
        value[key].json;
      }
    }

    return value;
  }

  get<K extends keyof T>(key: K): T[K] {
    return this.value[key];
  }

  set<K extends keyof T>(key: K, value: T[K]) {
    const v = this.value;
    v[key] = value;
    this.next(v);
  }

  del<K extends keyof T>(key: K) {
    const v = this.value;
    delete v[key];
    this.next(v);
  }
}
