import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { ThreeJSComponent } from "./threejs.component";
import { GameOverlayServiceModule } from "../overlay/service/game-overlay-service.module";

@NgModule({
    imports: [
        // angular
        BrowserModule,
        //custom
        GameOverlayServiceModule
    ],
    declarations: [
        ThreeJSComponent
    ],
    exports: [
        ThreeJSComponent
    ]
})
export class ThreeJSComponentModule {

}