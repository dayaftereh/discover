import * as THREE from 'three';
import { FlyControls } from 'three/examples/jsm/controls/FlyControls';
import { MethodCall } from '@angular/compiler';
import { Quaternion } from 'three';
import { GameComponent } from './game-component';

export class Player implements GameComponent {

    private object: THREE.Object3D

    constructor() {
        this.object = new THREE.Object3D()
    }

    gameObject(): THREE.Object3D {
        return this.object
    }

    init(): void {
        const geometry: THREE.SphereGeometry = new THREE.SphereGeometry(2.0)
        const material: THREE.MeshBasicMaterial = new THREE.MeshBasicMaterial(
            { color: 0xff0000, wireframe: true }
        )
        const mesh: THREE.Mesh = new THREE.Mesh(geometry, material)

        this.object.add(mesh)
        this.object.add(new THREE.AxesHelper(3))
    }

    update(delta: number): void {


    }

    dispose(): void {

    }

}
