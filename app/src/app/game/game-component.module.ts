import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { GameComponent } from "./game.component";
import { ThreeJSComponentModule } from "./threejs/threejs-component.module";
import { ConnectionServiceModule } from "../services/api/connection/connection-service.module";

@NgModule({
    imports: [
        // angular
        BrowserModule,
        // custom
        ThreeJSComponentModule,
        ConnectionServiceModule
    ],
    declarations: [
        GameComponent
    ]
})
export class GameComponentModule {

}
