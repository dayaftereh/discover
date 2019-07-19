import * as THREE from 'three';

export interface ObjectsInfo {
    player?: {
        position: THREE.Vector3,
        rotation: THREE.Euler
    },
    camera?: {
        position: THREE.Vector3,
        rotation: THREE.Euler
    },
    input?: {
        position: THREE.Vector3,
        rotation: THREE.Vector3
    }
}