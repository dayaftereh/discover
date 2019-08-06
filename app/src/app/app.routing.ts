import { RouterModule, Routes } from '@angular/router';
import { NgModule } from "@angular/core";
import { LoginComponent } from './login/login.component';
import { GameComponent } from './game/game.component';
import { GuardServicesModule } from './services/guard/guard-services.module';
import { InGameGuardService } from './services/guard/in-game-guard.service';
import { AuthenticationGuardService } from './services/guard/authentication-guard.service';
import { PlanetPreviewComponent } from './planet-preview/planet-preview.component';
import { AdminComponent } from './admin/admin.component';
import { AdminStarSystemsComponent } from './admin/star-systems/admin-star-systems.component';
import { AdminStarSystemComponent } from './admin/star-systems/admin-star-system.component';

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
    {
        path: 'admin',
        canActivate: [AuthenticationGuardService],
        children: [
            { path: '', component: AdminComponent },
            { path: 'star-system/:name', component: AdminStarSystemComponent },
            { path: 'star-systems', component: AdminStarSystemsComponent },
        ]
    },
    { path: 'planet-preview', component: PlanetPreviewComponent },
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