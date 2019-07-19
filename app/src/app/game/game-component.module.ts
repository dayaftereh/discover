import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { GameComponent } from "./game.component";
import { ThreeJSComponentModule } from "./threejs/threejs-component.module";
import { ConnectionServiceModule } from "../services/api/connection/connection-service.module";
import { GameServiceModule } from "../services/api/game/game-service.module";
import { GameOverlayComponentModule } from "./overlay/game-overlay-component.module";

@NgModule({
    imports: [
        // angular
        BrowserModule,
        // custom
        GameServiceModule,
        ThreeJSComponentModule,
        ConnectionServiceModule,
        GameOverlayComponentModule
    ],
    declarations: [
        GameComponent
    ]
})
export class GameComponentModule {

}
