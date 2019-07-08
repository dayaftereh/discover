import { RouterModule, Routes } from '@angular/router';
import { NgModule } from "@angular/core";
import { LoginComponent } from './login/login.component';
import { GameComponent } from './game/game.component';
import { GuardServicesModule } from './services/guard/guard-services.module';
import { InGameGuardService } from './services/guard/in-game-guard.service';
import { AuthenticationGuardService } from './services/guard/authentication-guard.service';

const routes: Routes = [
    {
        path: 'login',
        canActivate: [InGameGuardService],
        children: [
            { path: '', component: LoginComponent },
        ]
    },
    {
        path: 'game',
        canActivate: [AuthenticationGuardService],
        children: [
            { path: '', component: GameComponent },
        ]
    },
    { path: '', redirectTo: '/login', pathMatch: 'full' }
]

@NgModule({
    imports: [
        // Angular
        RouterModule.forRoot(routes),
        // Custom
        GuardServicesModule
    ]
})
export class AppRouting {

}