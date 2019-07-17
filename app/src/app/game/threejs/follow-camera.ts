import * as THREE from 'three';
import { Player } from './player';

export class FollowCamera {

    private offset: THREE.Vector3

    private damping: number = 2.5

    private cameraOrientation: THREE.Quaternion

    constructor(private readonly camera: THREE.Camera,
        private readonly player: Player, offset?: THREE.Vector3) {
        if (!offset) {
            offset = new THREE.Vector3(0, 0, -5)
        }
        this.offset = offset
        // rotation for camera
        this.cameraOrientation = new THREE.Quaternion()
        this.cameraOrientation.setFromEuler(new THREE.Euler(0, Math.PI, 0))
    }

    update(delta: number): void {
        if (!this.player || !this.player.object) {
            return
        }

        const offset: THREE.Vector3 = this.offset.clone()

        // calculate world point for offset
        const position: THREE.Vector3 = this.player.object.localToWorld(offset)

        // update the camera location
        this.camera.position.copy(position)

        // get the player rotation
        const playerRotation: THREE.Quaternion = this.player.rotation()
        // calculate target rotation
        const targetRotation: THREE.Quaternion = playerRotation.multiply(this.cameraOrientation)
        
        // get the current camera rotation
        const cameraRotation: THREE.Quaternion = this.camera.quaternion.clone()
        // slerp the rotation for daming
        const rotation: THREE.Quaternion = cameraRotation.slerp(targetRotation, delta * this.damping)

        // update the camera rotation
        this.camera.quaternion.copy(rotation)
    }

}