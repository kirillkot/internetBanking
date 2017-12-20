import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ConverterListComponent } from './converter-list.component';

const routes: Routes = [
    { path: 'converter', component: ConverterListComponent },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class ConverterRoutingModule { }
