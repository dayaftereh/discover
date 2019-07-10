import * as THREE from 'three';
import { FlyControls } from 'three/examples/jsm/controls/FlyControls';
import { MethodCall } from '@angular/compiler';

export class Player {

    mesh: THREE.Mesh | undefined

    geometry: THREE.SphereBufferGeometry | undefined

    material: THREE.MeshBasicMaterial | undefined

    constructor() {

    }

    init(): void {
        this.geometry = new THREE.SphereBufferGeometry(10.0)
        this.material = new THREE.MeshBasicMaterial()
        this.mesh = new THREE.Mesh(this.geometry, this.material)
    }

    update0(position: THREE.Vector3, rotation: THREE.Vector3) {
        if (this.mesh) {
            // set the mesh position
            this.mesh.position = position

            // get the euler from rotation
            const euler: THREE.Euler = new THREE.Euler()
            euler.setFromVector3(rotation)

            // set the rotation to quaternion
            this.mesh.quaternion.setFromEuler(euler)
        }
    }

    update(delta: number): void {

    }

    eulerRotation(): THREE.Euler {
        const rotation: THREE.Euler = new THREE.Euler()
        if (this.mesh) {
            rotation.setFromQuaternion(this.mesh.quaternion)
        }
        return rotation
    }

    position(): THREE.Vector3 {
        if (this.mesh) {
            return this.mesh.position
        }
        return new THREE.Vector3()
    }

    dispose(): void {

    }

}
