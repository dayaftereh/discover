import * as THREE from 'three';
import { FlyControls } from 'three/examples/jsm/controls/FlyControls';
import { MethodCall } from '@angular/compiler';
import { Quaternion } from 'three';
import { GameComponent } from './game-component';

export class Player implements GameComponent {

    private positon: THREE.Vector3
    private rotation: THREE.Quaternion

    constructor() {
        this.positon = new THREE.Vector3()
        this.rotation = new THREE.Quaternion()
    }

    init(): void {
    }

    playerPosition(): THREE.Vector3 {
        return this.positon.clone()
    }

    playerRotation(): THREE.Quaternion {
        return this.rotation.clone()
    }

    updatePositon(positon: THREE.Vector3): void {
        this.positon.copy(positon)
    }

    updateRotation(rotation: THREE.Vector3): void {
        const euler: THREE.Euler = new THREE.Euler(rotation.x, rotation.y, rotation.z)
        this.rotation.setFromEuler(euler)
    }

    update(delta: number): void {
    }

    dispose(): void {
    }

}
