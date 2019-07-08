import { AfterViewInit, Component, ElementRef, HostListener, ViewChild } from "@angular/core";
import * as THREE from 'three';

@Component({
    selector: 'app-threejs',
    templateUrl: './threejs.component.html'
})
export class ThreeJSComponent implements AfterViewInit {

    @ViewChild('canvas', { static: false })
    canvasRef: ElementRef | undefined;

    private scene: THREE.Scene | undefined;
    private renderer: THREE.WebGLRenderer | undefined
    private camera: THREE.PerspectiveCamera | undefined

    constructor() {

    }

    ngAfterViewInit(): void {
        this.createScene()
        this.createCamera()
        this.createRenderer()
    }

    private canvas(): HTMLCanvasElement {
        if (this.canvasRef) {
            return this.canvasRef.nativeElement;
        }
        throw new Error("unable to find canvas ref")
    }

    private calculateAspectRatio(): number {
        const canvas: HTMLCanvasElement = this.canvas()
        const height: number = canvas.clientHeight
        if (height > 0.0) {
            return canvas.clientWidth / height
        }
        return 0.0
    }

    private createScene(): void {
        this.scene = new THREE.Scene()
        this.scene.add(new THREE.AxesHelper(200))
    }

    private createCamera(): void {
        const aspectRatio: number = this.calculateAspectRatio()
        this.camera = new THREE.PerspectiveCamera(60, aspectRatio, 1, 1100)

        this.camera.position.x = 10
        this.camera.position.y = 10
        this.camera.position.z = 100
    }

    private createRenderer(): void {
        const canvas: HTMLCanvasElement = this.canvas()

        this.renderer = new THREE.WebGLRenderer({
            canvas,
            antialias: true,
        })

        this.renderer.setPixelRatio(devicePixelRatio)
        this.renderer.setSize(400, 400)
    }

    @HostListener('window:resize', ['$event'])
    public onResize(event: Event) {
        const canvas: HTMLCanvasElement = this.canvas()
        canvas.style.width = "100%";
        canvas.style.height = "100%";

        if (this.camera) {
            const aspectRatio: number = this.calculateAspectRatio()
            this.camera.aspect = aspectRatio;
            this.camera.updateProjectionMatrix();
        }

        if (this.renderer) {
            this.renderer.setSize(canvas.clientWidth, canvas.clientHeight)
        }

    }

    private update(): void {
        requestAnimationFrame(() => {
            this.update()
        })

        this.render()
    }

    private render(): void {
        if (this.renderer && this.scene && this.camera) {
            this.renderer.render(this.scene, this.camera)
        }
    }

}