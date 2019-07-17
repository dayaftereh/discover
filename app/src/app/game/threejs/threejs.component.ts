import { AfterViewInit, Component, ElementRef, EventEmitter, HostListener, OnDestroy, Output, ViewChild } from "@angular/core";
import * as THREE from 'three';
import { ThreeJSInitEvent } from "./threejs-init-event";
import { ThreeJSUpdateEvent } from "./threejs-update-event";
import { Subject } from "rxjs";

@Component({
    selector: 'app-threejs',
    templateUrl: './threejs.component.html'
})
export class ThreeJSComponent implements AfterViewInit, OnDestroy {

    @Output("onInit")
    onInit: Subject<ThreeJSInitEvent>

    @Output("onUpdate")
    onUpdate: Subject<ThreeJSUpdateEvent>

    @ViewChild('canvas', { static: false })
    canvasRef: ElementRef | undefined;

    private running: boolean

    private clock: THREE.Clock

    private scene: THREE.Scene | undefined;
    private renderer: THREE.WebGLRenderer | undefined
    private camera: THREE.PerspectiveCamera | undefined

    constructor() {
        this.running = false
        this.clock = new THREE.Clock()
        this.onInit = new Subject<ThreeJSInitEvent>()
        this.onUpdate = new Subject<ThreeJSUpdateEvent>()
    }

    ngAfterViewInit(): void {
        // create scene camera and renderer
        this.createScene()
        this.createCamera()
        this.createRenderer()

        // notify about successful init
        this.emitInit()

        // start the update loop
        this.start()
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
    }

    private createCamera(): void {
        const aspectRatio: number = this.calculateAspectRatio()
        this.camera = new THREE.PerspectiveCamera(60, aspectRatio, 1, 1100)
        this.camera.position.set(0, 0, 0)
    }

    private createRenderer(): void {
        const canvas: HTMLCanvasElement = this.canvas()

        this.renderer = new THREE.WebGLRenderer({
            canvas,
            antialias: true,
        })

        this.renderer.setPixelRatio(devicePixelRatio)
        this.renderer.setSize(1024, 800)
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

    private emitInit(): void {
        // get the canvas
        const canvas: HTMLCanvasElement = this.canvas()

        // fire the init event
        if (this.scene && this.renderer && this.camera) {
            const event: ThreeJSInitEvent = {
                canvas,
                scene: this.scene,
                camera: this.camera,
                renderer: this.renderer
            }
            this.onInit.next(event)
        }

        this.onInit.complete()
    }

    private start(): void {
        this.running = true
        this.update()
    }

    private update(): void {
        requestAnimationFrame(() => {
            if (this.running) {
                this.update()
            }
        })

        // get the delta for this loop
        const delta: number = this.clock.getDelta()

        // get the canvas
        const canvas: HTMLCanvasElement = this.canvas()

        // fire the update event
        if (this.scene && this.renderer && this.camera) {
            const event: ThreeJSUpdateEvent = {
                delta,
                canvas,
                scene: this.scene,
                camera: this.camera,
                renderer: this.renderer
            }

            this.onUpdate.next(event)
        }

        this.render()
    }

    private render(): void {
        if (this.renderer && this.scene && this.camera) {
            this.renderer.render(this.scene, this.camera)
        }
    }

    ngOnDestroy(): void {
        // disable update loop
        this.running = false

        // notify about end
        this.onUpdate.complete()

        // remove the renderer
        if (this.renderer) {
            this.renderer.dispose()
        }

        // remove the scene
        if (this.scene) {
            this.scene.dispose()
        }
    }

}