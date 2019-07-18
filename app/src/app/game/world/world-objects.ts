import * as THREE from 'three';
import { GameObject } from 'src/app/services/api/connection/messages/world-update';

export class WorldObjects {

    private objects: Map<number, THREE.Object3D>

    constructor(private readonly scene: THREE.Scene) {
        this.objects = new Map<number, THREE.Object3D>()
    }

    updateWorldObjects(objects: { [key: number]: GameObject }): void {
        // check if objects available
        if (!objects) {
            return
        }

        // get the keys
        const keys: string[] = Object.keys(objects)

        keys.forEach((identifier: string) => {
            // get the key of the object
            const key: number = Number(identifier)
            // get the game object
            const gameObject: GameObject = objects[key]
            // update the object
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
        // get the game object
        const object: THREE.Object3D | undefined = this.objects.get(key)
        if (object) {
            this.scene.remove(object)
        }
        // remove it from the world objects
        this.objects.delete(key)
    }

    private updateGameObject(key: number, gameObject: GameObject): void {
        const object: THREE.Object3D | undefined = this.objects.get(key)
        if (!object) {
            throw new Error(`unable to find game object [ ${key} ] in world`);
        }

        // update the position
        if (gameObject.position) {
            object.position.set(gameObject.position.x, gameObject.position.y, gameObject.position.z)
        }

        // update the rotation
        if (gameObject.rotation) {
            const euler: THREE.Euler = new THREE.Euler(gameObject.rotation.x, gameObject.rotation.y, gameObject.rotation.z)
            object.quaternion.setFromEuler(euler)
        }
    }

    private createGameObject(key: number, gameObject: GameObject): void {
        // get the radius
        let radius: number = 1
        if (gameObject.radius && gameObject.radius > 0) {
            radius = gameObject.radius
        }

        // get the color
        let color: number = 0xff0000
        if (gameObject.color && gameObject.color > 0) {
            color = gameObject.color
        }

        // create the three.js object
        const geometry: THREE.SphereGeometry = new THREE.SphereGeometry(radius)
        const material: THREE.MeshBasicMaterial = new THREE.MeshBasicMaterial(
            {
                color,
                wireframe: true,
            }
        )
        const mesh: THREE.Mesh = new THREE.Mesh(geometry, material)

        // create the THREE.Object3D
        const object: THREE.Object3D = new THREE.Object3D()
        object.add(new THREE.AxesHelper(radius + radius * .25))
        object.add(mesh)

        // add the object to world
        this.objects.set(key, object)

        // update the game object
        this.updateGameObject(key, gameObject)

        // add the object to scene
        this.scene.add(object)
    }

    getObjectById(id: number): THREE.Object3D | undefined {
        const object: THREE.Object3D | undefined = this.objects.get(id)
        return object
    }

}