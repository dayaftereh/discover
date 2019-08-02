import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { RouterModule } from '@angular/router';
import { LoginComponentModule } from "../login/login-component.module";
import { LayoutComponent } from "./layout.component";
import { GameComponentModule } from "../game/game-component.module";
import { NavbarComponentModule } from "../navbar/navbar-component.module";
import { PlanetPreviewComponentModule } from "../planet-preview/planet-preview-component.module";

@NgModule({
    imports: [
        // angular
        RouterModule,
        BrowserModule,
        // custom
        GameComponentModule,
        LoginComponentModule,
        NavbarComponentModule,
        PlanetPreviewComponentModule
    ],
    declarations: [
        LayoutComponent
    ],
    exports: [
        LayoutComponent
    ]
})
export class LayoutModule {

}