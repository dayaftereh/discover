import { Component, OnDestroy, OnInit } from "@angular/core";
import { ConnectionService } from "../services/api/connection/connection.service";
import { GameService } from "../services/api/game/game.service";
import { ThreeJSInitEvent } from "./threejs/threejs-init-event";
import { ThreeJSUpdateEvent } from "./threejs/threejs-update-event";
import { Game } from "./world/game";

@Component({
    templateUrl: './game.component.html'
})
export class GameComponent implements OnInit, OnDestroy {

    private game: Game

    constructor(private readonly connectionService: ConnectionService,
        private readonly gameService: GameService) {
        this.game = new Game(connectionService)
    }

    async ngOnInit(): Promise<void> {
        // handle server connection
        await this.connect()
        // notify server about ready
        await this.gameService.ready()
    }

    private async connect(): Promise<void> {
        // check if already connected
        const isConnected: boolean = this.connectionService.isConnected()
        if (isConnected) {
            return
        }
        // connect via websocket
        await this.connectionService.connect()
    }

    onInitThreeJS(event: ThreeJSInitEvent): void {
        this.game.init(event)
    }

    onUpdateThreeJS(event: ThreeJSUpdateEvent): void {        
        this.game.update(event)
    }

    ngOnDestroy(): void {
        this.game.dispose()
    }

}