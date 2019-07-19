import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { GameInfoOverlayComponent } from "./game-info-overlay.component";

@NgModule({
    imports: [
        BrowserModule
    ],
    declarations: [
        GameInfoOverlayComponent
    ],
    exports: [
        GameInfoOverlayComponent
    ]
})
export class GameInfoOverlayComponentModule {

}