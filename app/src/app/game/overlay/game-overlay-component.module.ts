import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { GameOverlayService } from "./service/game-overlay.service";
import { GameOverlayComponent } from "./game-overlay.component";
import { GameInfoOverlayComponentModule } from "./info/game-info-overlay-component.module";
import { GameOverlayServiceModule } from "./service/game-overlay-service.module";

@NgModule({
    imports: [
        // angular
        BrowserModule,
        //custom
        GameOverlayServiceModule,
        GameInfoOverlayComponentModule,
    ],
    declarations: [
        GameOverlayComponent
    ],
    exports: [
        GameOverlayComponent
    ]
})
export class GameOverlayComponentModule {

}