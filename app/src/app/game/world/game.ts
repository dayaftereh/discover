import { ThreeJSInitEvent } from "../threejs/threejs-init-event";
import { ThreeJSUpdateEvent } from "../threejs/threejs-update-event";
import { ConnectionService } from "src/app/services/api/connection/connection.service";
import { World } from "./world";
import { Player } from "./player";
import { Dispatcher } from "./dispatcher";
import { GameComponent } from "./game-component";
import { FollowCamera } from "./follow-camera";
import { GameInput } from "./game-input";
import { GameOverlayService } from "../overlay/service/game-overlay.service";

export class Game {

    private components: GameComponent[]

    constructor(private readonly connectionService: ConnectionService,
        private readonly gameOverlayService: GameOverlayService) {
        this.components = []
    }

    private createComponents(event: ThreeJSInitEvent): GameComponent[] {
        const player: Player = new Player(this.gameOverlayService)
        const followCamera: FollowCamera = new FollowCamera(event.camera, player, this.gameOverlayService)
        const world: World = new World(event.scene, player)

        const gameInput: GameInput = new GameInput(event.canvas, this.gameOverlayService, this.connectionService)
        const dispatcher: Dispatcher = new Dispatcher(world, this.connectionService)

        return [
            gameInput, player, followCamera, world, dispatcher
        ]
    }

    init(event: ThreeJSInitEvent): void {
        console.log("Init")
        this.components = this.createComponents(event)
        this.components.forEach((element: GameComponent) => {
            element.init()
        })
    }

    update(event: ThreeJSUpdateEvent): void {
        this.components.forEach((element: GameComponent) => {
            element.update(event.delta)
        })
    }

    dispose(): void {
        this.components.forEach((element: GameComponent) => {
            element.dispose()
        })
    }

}