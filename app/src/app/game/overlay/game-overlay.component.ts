import { Component } from "@angular/core";
import { GameOverlayService } from "./service/game-overlay.service";

@Component({
    selector: 'app-game-overlay',
    templateUrl: './game-overlay.component.html',
    styleUrls:[
        './game-overlay.component.scss'
    ]
})
export class GameOverlayComponent {

    constructor(private readonly gameOverlayService: GameOverlayService) {

    }


}