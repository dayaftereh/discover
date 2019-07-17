import * as THREE from 'three';
import { FlyControls } from 'three/examples/jsm/controls/FlyControls';
import { MethodCall } from '@angular/compiler';
import { Quaternion } from 'three';

export class Player {

    mesh: THREE.Mesh | undefined

    object: THREE.Object3D

    geometry: THREE.SphereGeometry | undefined

    material: THREE.MeshBasicMaterial | undefined

    constructor() {
        this.object = new THREE.Object3D()
    }

    init(): void {
        this.geometry = new THREE.SphereGeometry(2.0)
        this.material = new THREE.MeshBasicMaterial(
            { color: 0xff0000, wireframe: true }
        )
        this.mesh = new THREE.Mesh(this.geometry, this.material)

        //this.object.add(this.mesh)
        this.object.add(new THREE.AxesHelper(3))
    }

    update0(position: THREE.Vector3, rotation: THREE.Vector3) {
        if (this.object) {
            this.object.position.copy(position)
            this.object.rotation.setFromVector3(rotation)
        }
    }

    update(delta: number): void {
        

    }

    rotation(): THREE.Quaternion {
        if (this.object) {
            return this.object.quaternion.clone()
        }
        return new THREE.Quaternion()
    }

    eulerRotation(): THREE.Euler {
        const rotation: THREE.Euler = new THREE.Euler()
        if (this.object) {
            return this.object.rotation.clone()
        }
        return new THREE.Euler()
    }

    position(): THREE.Vector3 {
        if (this.object) {
            return this.object.position.clone()
        }
        return new THREE.Vector3()
    }

    dispose(): void {

    }

}
