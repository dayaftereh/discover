import { Component } from "@angular/core";
import * as THREE from 'three';
import { ThreeJSInitEvent } from "../game/threejs/threejs-init-event";
import { ThreeJSUpdateEvent } from "../game/threejs/threejs-update-event";
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls.js';
import { WebGLRenderer, Euler } from "three";


@Component({
    templateUrl: './planet-preview.component.html'
})
export class PlanetPreviewComponent {

    private light: THREE.PointLight | undefined
    private cloud: THREE.Mesh | undefined
    private cloudMaterial: THREE.MeshPhongMaterial | undefined
    private planetMaterial: THREE.MeshPhongMaterial | undefined

    private controls: OrbitControls | undefined

    constructor() {
    }

    onInitThreeJS(event: ThreeJSInitEvent): void {
        event.camera.position.set(3.0, 2.0, 0)

        const directionalLight: THREE.DirectionalLight = new THREE.DirectionalLight(0xffffff, 1);
        directionalLight.position.set(1, 5, 10).normalize();
        event.scene.add(directionalLight);

        const planet: THREE.Mesh = this.createPlanet()
        event.scene.add(planet)

        this.light = new THREE.PointLight(0xffffff, 1.0, 0.0, 1.0)
        this.light.position.set(-10.0, 0.0, 0.0)
        event.scene.add(this.light)

        this.controls = new OrbitControls(event.camera, event.canvas);
        this.controls.rotateSpeed = 0.1
        this.controls.zoomSpeed = 0.1
        this.controls.panSpeed = 0.1
        this.controls.autoRotate = true
        this.controls.autoRotateSpeed = 0.1
        this.controls.enableDamping = true;
        this.controls.dampingFactor = 0.05;
        this.controls.screenSpacePanning = false;
        this.controls.minDistance = 0.5;
        this.controls.maxDistance = 15;
    }

    private createPlanet(): THREE.Mesh {
        const radius: number = 1.0
        const geometry: THREE.SphereGeometry = new THREE.SphereGeometry(radius, 100, 100);

        this.planetMaterial = new THREE.MeshPhongMaterial();
        this.planetMaterial.bumpScale = 0.05
        this.planetMaterial.specular = new THREE.Color('gray')

        const mesh: THREE.Mesh = new THREE.Mesh(geometry, this.planetMaterial)

        const cloudGeometry: THREE.SphereGeometry = new THREE.SphereGeometry(radius * 1.05, 100, 100);
        this.cloudMaterial = new THREE.MeshPhongMaterial({
            side: THREE.DoubleSide,
            opacity: 1.0,
            transparent: true,
            depthWrite: false,
        })

        this.cloud = new THREE.Mesh(cloudGeometry, this.cloudMaterial)
        this.cloud.visible = false

        mesh.add(this.cloud)

        return mesh
    }

    async onUploaded(event: any): Promise<void> {
        if (!event || !event.files || event.files.length < 1 || !this.cloudMaterial || !this.planetMaterial) {
            return
        }

        const file: File = event.files[0]
        const allInOneTexture: THREE.Texture = await this.loadTexture(file)

        const textures: THREE.Texture[] = this.splitTexture(allInOneTexture, 5, 1)

        const biomeTexture: THREE.Texture = textures[0]
        const normalTexture: THREE.Texture = textures[1]
        const specularTexture: THREE.Texture = textures[2]
        const bumpTexture: THREE.Texture = textures[3]

        this.planetMaterial.map = biomeTexture
        this.planetMaterial.normalMap = normalTexture
        this.planetMaterial.specularMap = specularTexture
        this.planetMaterial.bumpMap = bumpTexture
        this.planetMaterial.needsUpdate = true

        const cloudTexture: THREE.Texture = textures[4]
        this.cloudMaterial.map = cloudTexture
        this.cloudMaterial.needsUpdate = true
    }

    private loadTexture(file: File): Promise<THREE.Texture> {
        const url: string = URL.createObjectURL(file)
        const loader: THREE.TextureLoader = new THREE.TextureLoader()

        return new Promise((resolve, reject) => {
            const texture: THREE.Texture = loader.load(`${url}`,
                () => {
                    resolve(texture)
                },
                undefined,
                (err: ErrorEvent) => {
                    reject(err)
                })
        })
    }

    private splitTexture(texture: THREE.Texture, x: number, y: number): THREE.Texture[] {
        const sourceCanvas: HTMLCanvasElement = document.createElement("canvas")
        sourceCanvas.width = texture.image.width
        sourceCanvas.height = texture.image.height

        const sourceContext: CanvasRenderingContext2D = sourceCanvas.getContext("2d") as CanvasRenderingContext2D
        sourceContext.drawImage(texture.image, 0, 0)

        const outputCanvasList: HTMLCanvasElement[] = []

        const xStep: number = sourceCanvas.width / x
        const yStep: number = sourceCanvas.height / y

        for (var i: number = 0; i < x; i++) {
            for (var j: number = 0; j < y; j++) {
                const offsetX: number = i * xStep
                const offsetY: number = j * yStep
                const imageData: ImageData = sourceContext.getImageData(offsetX, offsetY, xStep, yStep)

                const outputCanvas: HTMLCanvasElement = document.createElement("canvas")
                outputCanvas.width = xStep
                outputCanvas.height = yStep

                const outputContext: CanvasRenderingContext2D = outputCanvas.getContext("2d") as CanvasRenderingContext2D
                outputContext.putImageData(imageData, 0, 0)

                outputCanvasList.push(outputCanvas)
            }
        }

        return outputCanvasList.map((canvas: HTMLCanvasElement) => {
            const output: THREE.Texture = new THREE.CanvasTexture(canvas)
            return output
        })
    }

    private rotateLight(): void {
        if (!this.light) {
            return
        }
        const speed: number = Date.now() / 2000.0
        const radius: number = 10.0
        const lx: number = radius * Math.cos(speed)
        const lz: number = radius * Math.sin(speed)

        const ly: number = 5.0 + 5.0 * Math.sin(speed / 3.0)

        this.light.position.set(lx, ly, lz)
        this.light.lookAt(new THREE.Vector3(0.0, 0.0, 0.0))
    }

    onUpdateThreeJS(event: ThreeJSUpdateEvent): void {
        if (this.controls) {
            this.controls.update()
        }
        if (this.light) {
            this.rotateLight()
        }
    }

    onCloudStateChanged(): void {
        if (this.cloud) {
            this.cloud.visible = !this.cloud.visible
        }
    }
}
