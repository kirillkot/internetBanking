import { NgModule } from '@angular/core';

import { ReactiveFormsModule } from '@angular/forms';

import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatRadioModule } from '@angular/material/radio';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatButtonModule } from '@angular/material/button';
import { MatListModule } from '@angular/material/list';
import { MatTableModule } from '@angular/material/table';

import { MatNativeDateModule } from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';


const modules = [
  ReactiveFormsModule,

  MatInputModule,
  MatSelectModule,
  MatCheckboxModule,
  MatFormFieldModule,
  MatRadioModule,
  MatDatepickerModule,
  MatButtonModule,
  MatListModule,
  MatTableModule,

  MatNativeDateModule,
  BrowserAnimationsModule
]


@NgModule({
  imports: modules,
  exports: modules,
  declarations: []
})
export class MaterialModule { }
