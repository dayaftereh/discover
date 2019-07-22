import { Subscription } from "rxjs";
import { ConnectionService } from "src/app/services/api/connection/connection.service";
import { Message } from "src/app/services/api/connection/messages/message";
import { MessageType } from "src/app/services/api/connection/messages/message-type";
import { WorldUpdate } from "src/app/services/api/connection/messages/world-update";
import { GameComponent } from "./game-component";
import { World } from "./world";
import { Pong } from "src/app/services/api/connection/messages/pong";
import { GameOverlayService } from "../overlay/service/game-overlay.service";

export class Dispatcher implements GameComponent {

    private subscription: Subscription | undefined

    constructor(private readonly world: World,
        private readonly connectionService: ConnectionService,
        private readonly gameOverlayService: GameOverlayService) {

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
            case MessageType.PONG:
                this.handlePong(message as Pong)
                break
        }
    }

    private handleWorldUpdate(update: WorldUpdate): void {
        // forword the world update
        this.world.worldUpdate(update)
    }

    private handlePong(pong: Pong): void {
        const now: number = Date.now()

        const clientDelta: number = now - pong.clientSendTime
        const serverDelta: number = pong.serverSendTime - pong.serverReceiveTime

        const delta: number = (clientDelta - serverDelta) / 2.0

        this.gameOverlayService.onEngineStats.emit({
            latency: delta
        })
    }

    update(delta: number): void { }

    dispose(): void { }
}