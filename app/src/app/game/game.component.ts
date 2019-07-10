import { Component, OnInit, OnDestroy, ViewChild, AfterViewInit } from "@angular/core";
import { ConnectionService } from "../services/api/connection/connection.service";
import { Movement } from "../services/api/connection/messages/movement";
import { Subscription } from "rxjs";
import { Message } from "../services/api/connection/messages/message";
import { MessageType } from "../services/api/connection/messages/message-type";
import { Update, GameObject } from "../services/api/connection/messages/update";
import { ThreeJSComponent } from "./threejs/threejs.component";
import * as THREE from 'three';
import { GameService } from "../services/api/game/game.service";

@Component({
    templateUrl: './game.component.html'
})
export class GameComponent implements OnInit, OnDestroy {

    @ViewChild("threeJS", { static: true })
    threeJS: ThreeJSComponent | undefined

    private ready: boolean

    private subscription: Subscription | undefined

    constructor(private readonly connectionService: ConnectionService,
        private readonly gameService: GameService) {
        this.ready = false
    }

    async  ngOnInit(): Promise<void> {
        // handle server connection
        await this.connect()
        // notify server about ready
        await this.gameService.ready()
        // mark it as ready
        this.ready = true
    }

    private async connect(): Promise<void> {
        // check if already connected
        const isConnected: boolean = this.connectionService.isConnected()
        if (isConnected) {
            return
        }
        // connect via websocket
        await this.connectionService.connect()

        // connect the dispatcher
        this.registerDispatcher()
    }

    private registerDispatcher(): void {
        // register for the on message
        this.subscription = this.connectionService.onMessage((message: Message) => {
            this.dispatch(message)
        })
    }

    private dispatch(message: Message): void {
        // check if the message is an update
        if (message.type !== MessageType.UPDATE) {
            return
        }
        // convert the message to the update
        const update: Update = message as Update;

        // update the player if exists
        if (update.player) {
            this.updatePlayerObject(update.player)
        }
    }

    private updatePlayerObject(playerObject: GameObject): void {
        if (playerObject.position && playerObject.rotation) {
            const position: THREE.Vector3 = new THREE.Vector3(
                playerObject.position.x,
                playerObject.position.y,
                playerObject.position.z
            )

            const rotation: THREE.Vector3 = new THREE.Vector3(
                playerObject.rotation.x,
                playerObject.rotation.y,
                playerObject.rotation.z
            )

            if (this.threeJS) {
                this.threeJS.updatePlayer(position, rotation)
            }
        }
    }

    onMovement(movement: Movement): void {
        if (this.ready) {
            this.connectionService.send(movement)
        }
    }

    ngOnDestroy(): void {
        this.ready = false;
        if (this.subscription) {
            this.subscription.unsubscribe()
        }
    }

}