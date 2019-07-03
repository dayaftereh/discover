import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { RouterModule } from '@angular/router';
import { LoginComponentModule } from "../login/login-component.module";
import { LayoutComponent } from "./layout.component";
import { GameComponentModule } from "../game/game-component.module";

@NgModule({
    imports: [
        // angular
        RouterModule,
        BrowserModule,
        // custom
        LoginComponentModule,
        GameComponentModule
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