import { Player } from "./player";
import * as THREE from 'three';
import { GameComponent } from "./game-component";
import { GameObject } from "src/app/services/api/connection/messages/update";

export class World implements GameComponent {

    private tick: number
    private objects: Map<number, THREE.Object3D>

    constructor(private readonly scene: THREE.Scene, private readonly player: Player) {
        this.tick = -1
        this.objects = new Map<number, THREE.Object3D>()
    }

    init(): void {
        this.scene.add(new THREE.AxesHelper(100))
    }

    worldUpdate(tick: number, player: GameObject, objects: { [key: number]: GameObject }): void {
        // check if world tick
        if (!(this.tick < tick)) {
            return
        }
        // update tick
        this.tick = tick;

        // update the player
        this.updatePlayerObject(player)

        // update the objects
        const keys: number[] = Object.keys(objects).map((key: string) => Number(key))
        keys.forEach((key: number) => {
            const gameObject: GameObject = objects[key]
            this.updateObject(key, gameObject)
        })
    }

    private updateObject(key: number, gameObject: GameObject): void {
        // check if the object needs to be removed
        if (gameObject.removeable) {
            this.removeGameObject(key)
            return
        }

        // only needs to update the object
        if (this.objects.has(key)) {
            this.updateGameObject(key, gameObject)
            return
        }

        // create a new game object
        this.createGameObject(key, gameObject)
    }

    private removeGameObject(key: number): void {
        const object: THREE.Object3D | undefined = this.objects.get(key)
        if (object) {
            this.scene.remove(object)
        }
        this.objects.delete(key)
    }

    private updateGameObject(key: number, gameObject: GameObject): void {
        const object: THREE.Object3D | undefined = this.objects.get(key)
        if (!object) {
            return
        }

        if (gameObject.position) {
            object.position.set(gameObject.position.x, gameObject.position.y, gameObject.position.z)
        }

        if (gameObject.rotation) {
            object.quaternion.setFromEuler(new THREE.Euler(gameObject.rotation.x, gameObject.rotation.y, gameObject.rotation.z))
        }
    }

    private updatePlayerObject(gameObject: GameObject): void {

        if (gameObject.position) {
            this.player.updatePositon(new THREE.Vector3(
                gameObject.position.x,
                gameObject.position.y,
                gameObject.position.z
            ))
        }

        if (gameObject.rotation) {
            this.player.updateRotation(new THREE.Vector3(
                gameObject.rotation.x,
                gameObject.rotation.y,
                gameObject.rotation.z
            ))
        }
    }

    private createGameObject(key: number, gameObject: GameObject): void {
        // get the radius
        let radius: number = 1
        if (gameObject.radius && gameObject.radius > 0) {
            radius = gameObject.radius
        }

        // create the game object
        const geometry: THREE.SphereGeometry = new THREE.SphereGeometry(radius)
        const material: THREE.MeshBasicMaterial = new THREE.MeshBasicMaterial(
            { color: 0xff0000, wireframe: true }
        )
        const mesh: THREE.Mesh = new THREE.Mesh(geometry, material)

        const object: THREE.Object3D = new THREE.Object3D()
        object.add(new THREE.AxesHelper(radius + radius * .25))
        object.add(mesh)

        // add the object
        this.objects.set(key, object)

        // update the obhect
        this.updateGameObject(key, gameObject)

        this.scene.add(object)
    }


    update(delta: number): void {
    }

    dispose(): void { }

}