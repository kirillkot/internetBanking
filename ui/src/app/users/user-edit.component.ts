import { Location } from '@angular/common'
import { Component } from '@angular/core'
import { FormBuilder, Validators } from '@angular/forms'

import { FormComponent } from '../abstract/form.component'

import { UserForm, User, UserService } from './user.service'

@Component({
    selector: 'app-edit-user',
    templateUrl: './user-edit.component.html',
    styleUrls: ['./user-edit.component.css']
})
export class UserEditComponent extends FormComponent<UserForm, User>{
    constructor(
        location: Location,
        formbuilder: FormBuilder,
        service: UserService,
    ) {
        super(location, formbuilder, service)
    }

    fields(): any {
      return {
          username: ['', [Validators.required, Validators.minLength(4)]],
          first_name: ['', [Validators.required, Validators.minLength(2)]],
          last_name: ['', [Validators.required, Validators.minLength(2)]],
          city_name: ['', [Validators.required, Validators.minLength(2)]],
          adress: ['', [Validators.required, Validators.minLength(2)]],
          password: ['', [Validators.required, Validators.minLength(4)]],
      };
    }
}
