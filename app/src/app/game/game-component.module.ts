import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { GameComponent } from "./game.component";
import { ThreeJSComponentModule } from "./threejs/threejs-component.module";

@NgModule({
    imports: [
        // angular
        BrowserModule,
        // custom
        ThreeJSComponentModule
    ],
    declarations: [
        GameComponent
    ]
})
export class GameComponentModule {

}