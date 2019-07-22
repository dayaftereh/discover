import * as THREE from 'three';
import { GameOverlayService } from '../overlay/service/game-overlay.service';
import { GameComponent } from './game-component';

export class Player implements GameComponent {

    private position: THREE.Vector3
    private rotation: THREE.Quaternion

    constructor(private readonly gameOverlayService: GameOverlayService) {
        this.position = new THREE.Vector3()
        this.rotation = new THREE.Quaternion()
    }

    init(): void {
    }

    playerPosition(): THREE.Vector3 {
        return this.position.clone()
    }

    playerRotation(): THREE.Quaternion {
        return this.rotation.clone()
    }

    updatePositon(position: THREE.Vector3): void {
        this.position.copy(position)
    }

    updateRotation(rotation: THREE.Euler): void {
        this.rotation.setFromEuler(rotation)
    }

    update(delta: number): void {
        // caluclate euler rotation
        const euler: THREE.Euler = new THREE.Euler()
        euler.setFromQuaternion(this.rotation)

        // notify overlay about it
        this.gameOverlayService.onObjectsInfo.emit({
            player: {
                rotation: euler,
                position: this.position.clone(),
            }
        })
    }

    dispose(): void {
    }

}
