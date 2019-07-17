import * as THREE from 'three';

export interface ThreeJSInitEvent {
    scene: THREE.Scene
    camera: THREE.PerspectiveCamera
    canvas: HTMLCanvasElement
    renderer: THREE.WebGLRenderer
}