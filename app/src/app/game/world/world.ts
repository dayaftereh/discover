import { Player } from "./player";
import * as THREE from 'three';
import { GameComponent } from "./game-component";
import { GameObject, WorldUpdate, PlayerUpdate } from "src/app/services/api/connection/messages/world-update";
import { WorldObjects } from "./world-objects";

export class World implements GameComponent {

    private tick: number
    private worldObjects: WorldObjects

    constructor(private readonly scene: THREE.Scene, private readonly player: Player) {
        this.tick = -1
        this.worldObjects = new WorldObjects(scene)
    }

    init(): void {
        this.scene.add(new THREE.AxesHelper(100))
    }

    worldUpdate(message: WorldUpdate): void {
        // check if this update is required
        if (!this.isUpdateRequired(message)) {
            return
        }

        // set the current server tick
        this.tick = message.tick

        // update the world objects
        this.worldObjects.updateWorldObjects(message.objects)

        // get the player update
        const player: PlayerUpdate = message.player

        // update the player
        this.updatePlayerObject(player)

    }

    private isUpdateRequired(message: WorldUpdate): boolean {
        if (this.tick < message.tick) {
            return true
        }
        return false
    }

    private updatePlayerObject(player: PlayerUpdate) {
        // get the player object    
        const object: THREE.Object3D | undefined = this.worldObjects.getObjectById(player.gameObjectId)
        // check if a object is found
        if (object) {
            // get the current player location
            const position: THREE.Vector3 = object.position.clone()
            // update the position
            this.player.updatePositon(position)

            // get the current player location
            const rotation: THREE.Euler = object.rotation.clone()
            // update the position
            this.player.updateRotation(rotation)
        }
    }

    update(delta: number): void {
    }

    dispose(): void { }

}