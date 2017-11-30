import { NgModule }              from '@angular/core';
import { RouterModule, Routes }  from '@angular/router';

import { BankComponent }         from './bank/bank.component';
import { PageNotFoundComponent } from './not-found/not-found.component';


const appRoutes: Routes = [
    { path: 'bank', component: BankComponent },
    { path: '',   redirectTo: '/bank', pathMatch: 'full' },
    { path: '**', component: PageNotFoundComponent }
];

@NgModule({
    imports: [
        RouterModule.forRoot(
            appRoutes,
            { enableTracing: true } // <-- debugging purposes only
        )
    ],
    exports: [
        RouterModule
    ]
})
export class AppRoutingModule { }
