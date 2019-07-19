import { AfterViewInit, Component, ElementRef, EventEmitter, HostListener, OnDestroy, Output, ViewChild, OnInit } from "@angular/core";
import * as THREE from 'three';
import { ThreeJSInitEvent } from "./threejs-init-event";
import { ThreeJSUpdateEvent } from "./threejs-update-event";
import { Subject, Subscription } from "rxjs";
import { Game } from "../world/game";
import { ConnectionService } from "src/app/services/api/connection/connection.service";
import { GameOverlayService } from "../overlay/service/game-overlay.service";
import { ThreeJSStats } from "./threejs-stats";

@Component({
    selector: 'app-threejs',
    templateUrl: './threejs.component.html',
    styleUrls: [
        './threejs.component.scss'
    ]
})
export class ThreeJSComponent implements OnInit, AfterViewInit, OnDestroy {

    @Output("onInit")
    onInit: EventEmitter<ThreeJSInitEvent>

    @Output("onUpdate")
    onUpdate: EventEmitter<ThreeJSUpdateEvent>

    @ViewChild('canvas', { static: false })
    canvasRef: ElementRef | undefined;

    private running: boolean

    private stats: ThreeJSStats

    private clock: THREE.Clock

    private scene: THREE.Scene | undefined;
    private renderer: THREE.WebGLRenderer | undefined
    private camera: THREE.PerspectiveCamera | undefined

    private subscription: Subscription | undefined

    constructor(private readonly gameOverlayService: GameOverlayService) {
        this.running = false
        this.clock = new THREE.Clock()
        this.stats = new ThreeJSStats()
        this.onInit = new EventEmitter<ThreeJSInitEvent>(true)
        this.onUpdate = new EventEmitter<ThreeJSUpdateEvent>(false)
    }

    ngOnInit(): void {
        //register for threejs stats events
        this.subscription = this.stats.onChanged.subscribe((event: { fps: number, update: number }) => {
            this.gameOverlayService.updateEngineStats(event.fps, 0, event.update)
        })
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
        // get the canvas
        const canvas: HTMLCanvasElement = this.canvas()

        // create the WebGLRenderer
        this.renderer = new THREE.WebGLRenderer({
            canvas,
            antialias: true,
        })

        // set the ratio
        this.renderer.setPixelRatio(devicePixelRatio)

        // set the size of the canvas
        const width: number = canvas.clientWidth
        const height: number = canvas.clientHeight
        this.renderer.setSize(width, height, true)
    }

    @HostListener('window:resize', ['$event'])
    public onResize(event: Event) {
        const canvas: HTMLCanvasElement = this.canvas()


        if (this.camera) {
            const aspectRatio: number = this.calculateAspectRatio()
            this.camera.aspect = aspectRatio;
            this.camera.updateProjectionMatrix();
        }

        if (this.renderer) {
            const width: number = canvas.clientWidth
            const height: number = canvas.clientHeight
            this.renderer.setSize(width, height, true)
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
        // start rendering
        this.running = true
        // start the lock
        this.clock.start()
        // execute first update
        this.update()
    }

    private update(): void {
        // request next frame
        requestAnimationFrame(() => {
            if (this.running) {
                this.update()
            }
        })
        // start the update
        this.stats.beginUpdate()

        // get the delta for this loop in seconds
        const delta: number = this.clock.getDelta()

        // get the canvas
        const canvas: HTMLCanvasElement = this.canvas()
        // fire the update event
        if (this.scene && this.renderer && this.camera) {
            // create update event
            const event: ThreeJSUpdateEvent = {
                delta,
                canvas,
                scene: this.scene,
                camera: this.camera,
                renderer: this.renderer
            }

            // fire the update event
            this.onUpdate.next(event)
        }

        // render the scene
        this.render()

        // end the update
        this.stats.endUpdate()
    }

    private render(): void {
        // render the scene with camera
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

        if (this.subscription) {
            this.subscription.unsubscribe()
        }
    }

}