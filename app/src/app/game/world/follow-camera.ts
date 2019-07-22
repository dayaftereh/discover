import * as THREE from 'three';
import { GameComponent } from './game-component';
import { Player } from './player';
import { GameOverlayService } from '../overlay/service/game-overlay.service';

export class FollowCamera implements GameComponent {

    private offset: THREE.Vector3

    private damping: number = 10.0

    private cameraOrientation: THREE.Quaternion

    constructor(private readonly camera: THREE.Camera,
        private readonly player: Player,
        private readonly gameOverlayService: GameOverlayService,
        offset?: THREE.Vector3) {
        if (!offset) {
            offset = new THREE.Vector3(0, 0, -5)
        }
        this.offset = offset

        // Up-Vector (1,0,0)
        // Move (0,0,1)
        this.camera.up.set(1, 0, 0)

        // rotation for camera
        this.cameraOrientation = new THREE.Quaternion()
    }

    init(): void {
        this.cameraOrientation.setFromEuler(new THREE.Euler(0, -Math.PI, Math.PI/2.0))
    }

    update(delta: number): void {
        const offset: THREE.Vector3 = this.offset.clone()

        const playerPosition: THREE.Vector3 = this.player.playerPosition()

        // get the player rotation
        const playerRotation: THREE.Quaternion = this.player.playerRotation()

        // calculate world point for offset
        const position: THREE.Vector3 = playerPosition.clone().add(offset.applyQuaternion(playerRotation.clone()))
        // update the camera location
        this.camera.position.copy(position)

        // calculate target rotation
        const targetRotation: THREE.Quaternion = playerRotation.multiply(this.cameraOrientation)
        // get the current camera rotation
        //const cameraRotation: THREE.Quaternion = this.camera.quaternion.clone()

        // clamp the damping between 0.0 and 1.0
        //const damping: number = this.clamp(delta * this.damping, 0.0, 1.0)
        
        // slerp the rotation for daming
        //const rotation: THREE.Quaternion = cameraRotation.rotateTowards(targetRotation, damping)

        // update the camera rotation
        this.camera.quaternion.copy(targetRotation)

        // notify overlay about new location and rotation
        this.gameOverlayService.onObjectsInfo.emit({
            camera: {
                position: this.camera.position.clone(),
                rotation: this.camera.rotation.clone(),
            }
        })
    }

    private clamp(v: number, min: number, max: number): number {
        return Math.min(Math.max(v, min), max)
    }

    dispose(): void {

    }
}