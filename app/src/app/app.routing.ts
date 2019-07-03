import { RouterModule, Routes } from '@angular/router';
import { NgModule } from "@angular/core";
import { LoginComponent } from './login/login.component';
import { GameComponent } from './game/game.component';

const routes: Routes = [
    { path: 'login', component: LoginComponent },
    { path: 'game', component: GameComponent },
    { path: '', redirectTo: '/login', pathMatch: 'full' }
]

@NgModule({
    imports: [
        // Angular
        RouterModule.forRoot(routes)
    ]
})
export class AppRouting {

}