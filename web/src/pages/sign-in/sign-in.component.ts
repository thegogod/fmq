import { ChangeDetectionStrategy, Component, OnInit, ViewEncapsulation } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';

import { Api } from '../../api';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrl: './sign-in.component.scss',
  host: { class: 'app-sign-in' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class SignInComponent implements OnInit {
  readonly form = new FormGroup({
    name: new FormControl<string>(''),
    password: new FormControl<string>('')
  });

  constructor(
    private readonly _router: Router,
    private readonly _api: Api
  ) { }

  ngOnInit() {
    document.title = 'FMQ Admin | Sign In';
  }

  async submit() {
    if (
      !this.form.value.name ||
      !this.form.value.password
    ) return;

    try {
      const token = btoa(`${this.form.value.name}:${this.form.value.password}`);
      localStorage.setItem('fmq--token', token);
      this._router.navigate(['/']);
    } catch (err) {
      throw err;
    }
  }
}
