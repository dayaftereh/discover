import * as THREE from 'three';
import { Player } from './player';
import { GameComponent } from './game-component';

export class FollowCamera implements GameComponent {

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
    }

    init(): void {
        this.cameraOrientation.setFromEuler(new THREE.Euler(0, Math.PI, 0))
    }

    update(delta: number): void {
        const offset: THREE.Vector3 = this.offset.clone()

        // get the player object
        const playerObject: THREE.Object3D = this.player.gameObject()

        // calculate world point for offset
        const position: THREE.Vector3 = playerObject.localToWorld(offset)

        // update the camera location
        this.camera.position.copy(position)

        // get the player rotation
        const playerRotation: THREE.Quaternion = playerObject.quaternion.clone()
        // calculate target rotation
        const targetRotation: THREE.Quaternion = playerRotation.multiply(this.cameraOrientation)

        // get the current camera rotation
        const cameraRotation: THREE.Quaternion = this.camera.quaternion.clone()
        // slerp the rotation for daming
        const rotation: THREE.Quaternion = cameraRotation.slerp(targetRotation, delta * this.damping)

        // update the camera rotation
        this.camera.quaternion.copy(rotation)
    }

    dispose(): void {

    }
}