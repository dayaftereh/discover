import { Subscription } from "rxjs";
import { ConnectionService } from "src/app/services/api/connection/connection.service";
import { Message } from "src/app/services/api/connection/messages/message";
import { MessageType } from "src/app/services/api/connection/messages/message-type";
import { WorldUpdate } from "src/app/services/api/connection/messages/world-update";
import { GameComponent } from "./game-component";
import { World } from "./world";

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
            case MessageType.WORLD_UPDATE:
                this.handleWorldUpdate(message as WorldUpdate)
                break
        }
    }

    private handleWorldUpdate(update: WorldUpdate): void {
        // forword the world update
        this.world.worldUpdate(update)
    }


    update(delta: number): void { }

    dispose(): void { }
}