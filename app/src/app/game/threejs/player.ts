import * as THREE from 'three';
import { FlyControls } from 'three/examples/jsm/controls/FlyControls';
import { MethodCall } from '@angular/compiler';

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

        this.object.add(this.mesh)
        this.object.add(new THREE.AxesHelper(3))
    }

    update0(position: THREE.Vector3, rotation: THREE.Vector3) {
        if (this.object) {
            //console.log("server", rotation)
            // set the mesh position
            this.object.position.copy(position)

            // get the euler from rotation
            //const euler: THREE.Euler = new THREE.Euler()
            //euler.setFromVector3(rotation)

            //const quaternion: THREE.Quaternion = new THREE.Quaternion()
            //quaternion.setFromEuler(euler)

            //this.mesh.setRotationFromEuler(euler)

            this.object.rotation.setFromVector3(rotation)
            //this.object.setRotationFromEuler(new THREE.Euler(0, 1, 0))
            //this.mesh.updateMatrix()
            //this.mesh.updateMatrixWorld()

            // set the rotation to quaternion
            //this.mesh.quaternion.copy(quaternion)
        }
    }

    update(delta: number): void {
        if (this.object) {
            //this.object.setRotationFromEuler(new THREE.Euler(0, 0.01, 0))
        }
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
