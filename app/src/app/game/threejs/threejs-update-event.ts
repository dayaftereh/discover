import * as THREE from 'three';

export interface ThreeJSUpdateEvent {
    delta: number
    scene: THREE.Scene
    camera: THREE.PerspectiveCamera
    canvas: HTMLCanvasElement
    renderer: THREE.WebGLRenderer
}