import { Component, OnInit } from "@angular/core";
import { ConnectionService } from "../services/api/connection/connection.service";

@Component({
    templateUrl: './game.component.html'
})
export class GameComponent implements OnInit {

    constructor(private readonly connectionService: ConnectionService) {

    }

    async  ngOnInit(): Promise<void> {
        // handle server connection
        await this.connect()
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

}