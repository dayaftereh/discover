import { Subscription } from "rxjs";
import { ConnectionService } from "src/app/services/api/connection/connection.service";
import { World } from "./world";
import { Message } from "src/app/services/api/connection/messages/message";
import { MessageType } from "src/app/services/api/connection/messages/message-type";
import { Update, GameObject } from "src/app/services/api/connection/messages/update";
import { Player } from "./player";
import { GameComponent } from "./game-component";
import * as THREE from 'three';

export class Dispatcher implements GameComponent {

    private subscription: Subscription | undefined

    constructor(private readonly world: World,
        private readonly connectionService: ConnectionService) {

    }

    init(): void {
        this.subscription = this.connectionService.onMessage((message: Message) => {
            this.dispatch(message)
        })
    }

    private dispatch(message: Message): void {
        switch (message.type) {
            case MessageType.UPDATE:
                this.handleUpdate(message as Update)
                break
        }
    }

    private handleUpdate(update: Update): void {        
        // forword the world update
        this.world.worldUpdate(update.tick, update.player, update.objects)
    }


    update(delta: number): void { }

    dispose(): void { }
}