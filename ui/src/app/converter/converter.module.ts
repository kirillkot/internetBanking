import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms'

import { MaterialModule } from '../material/material.module';

import { ConverterRoutingModule } from './converter-routing.module';
import { ConverterListComponent } from './converter-list.component';
import { ConverterService } from './converter.service';

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
        MaterialModule,
        ConverterRoutingModule,
        FormsModule,
    ],
    declarations: [
        ConverterListComponent,
    ],
    providers: [
        ConverterService,
    ]
})

export class ConverterModule { }
