import * as THREE from 'three';
import { Player } from './player';

export class FollowCamera {

    private offset: THREE.Vector3

    constructor(private readonly camera: THREE.Camera,
        private readonly player: Player, offset?: THREE.Vector3) {
        if (!offset) {
            offset = new THREE.Vector3(0, 0, 1)
        }
        this.offset = offset
    }

    update(delta: number): void {
        const targetRotation: THREE.Euler = this.player.eulerRotation()

        // only intresstet on y
        targetRotation.x = 0.0
        targetRotation.z = 0.0

        // create the rotation
        const rotaion: THREE.Quaternion = new THREE.Quaternion()
        rotaion.setFromEuler(targetRotation)

        // get the player position
        const position: THREE.Vector3 = this.player.position()
        // calculate camera location
        const cameraPosition: THREE.Vector3 = position.sub(this.offset.applyQuaternion(rotaion))

        // set the camera to the position
        this.camera.position.set(cameraPosition.x, cameraPosition.y, cameraPosition.z)

        // look at the position
        this.camera.lookAt(position)
    }

}