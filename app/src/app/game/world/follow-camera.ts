import * as THREE from 'three';
import { Player } from './player';
import { GameComponent } from './game-component';

export class FollowCamera implements GameComponent {

    private offset: THREE.Vector3

    private damping: number = 90.0

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

        const playerPosition: THREE.Vector3 = this.player.playerPosition()

        // get the player rotation
        const playerRotation: THREE.Quaternion = this.player.playerRotation()

        // calculate world point for offset
        const position: THREE.Vector3 = playerPosition.add(offset.applyQuaternion(playerRotation.clone()))
        // update the camera location
        this.camera.position.copy(position)

        // calculate target rotation
        const targetRotation: THREE.Quaternion = playerRotation.multiply(this.cameraOrientation)
        // get the current camera rotation
        const cameraRotation: THREE.Quaternion = this.camera.quaternion.clone()
        // slerp the rotation for daming
        const rotation: THREE.Quaternion = cameraRotation.slerp(targetRotation, delta * this.damping)

        // update the camera rotation
        this.camera.quaternion.copy(targetRotation)
    }

    dispose(): void {

    }
}