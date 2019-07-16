import * as THREE from 'three';
import { Player } from './player';

export class FollowCamera {

    private offset: THREE.Vector3

    count = 0

    constructor(private readonly camera: THREE.Camera,
        private readonly player: Player, offset?: THREE.Vector3) {
        if (!offset) {
            offset = new THREE.Vector3(0, 0, 15)
        }
        this.offset = offset

    }

    update(delta: number): void {
        const targetRotation: THREE.Euler = this.player.eulerRotation()

        // only intresstet on y
        //targetRotation.x = 0.0
        //targetRotation.z = 0.0

        // get the player position
        const position: THREE.Vector3 = this.player.position()
        // calculate camera location

        const offset: THREE.Vector3 = this.offset.clone()
        const cameraPosition: THREE.Vector3 = position.clone().sub(offset.applyEuler(targetRotation))

        // set the camera to the position
        this.camera.position.copy(cameraPosition)

        // look at the position
        this.camera.lookAt(position)
    }

}