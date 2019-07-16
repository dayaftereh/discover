import { AfterViewInit, Component, ElementRef, HostListener, ViewChild, EventEmitter, Output, OnDestroy, OnInit } from "@angular/core";
import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';
import { FlyControls } from "./fly-controls";
import { Player } from "./player";
import { FollowCamera } from "./follow-camera";
import { Movement } from "../../services/api/connection/messages/movement";
import { Subscription } from "rxjs";

@Component({
    selector: 'app-threejs',
    templateUrl: './threejs.component.html'
})
export class ThreeJSComponent implements OnInit, AfterViewInit, OnDestroy {

    @Output("onMovement")
    onMovement: EventEmitter<Movement> | undefined

    @ViewChild('canvas', { static: false })
    canvasRef: ElementRef | undefined;

    private running: boolean

    private clock: THREE.Clock

    private scene: THREE.Scene | undefined;
    private renderer: THREE.WebGLRenderer | undefined
    private camera: THREE.PerspectiveCamera | undefined

    private player: Player | undefined
    private controls: FlyControls | undefined
    private followCamera: FollowCamera | undefined

    private subscription: Subscription | undefined

    stats: any = {
        cameraLocation: new THREE.Vector3(),
        cameraRotation: new THREE.Euler(),
        playerLocation: new THREE.Vector3(),
        playerRotation: new THREE.Euler(),
        playerQuaternion: new THREE.Quaternion(),
    }


    constructor() {
        this.running = true
        this.clock = new THREE.Clock()
        this.onMovement = new EventEmitter<Movement>(true)
    }

    ngOnInit(): void {

    }


    ngAfterViewInit(): void {
        this.createScene()
        this.createCamera()
        this.createRenderer()

        this.createPlayer()
        this.createControls()
        this.createFollowCamera()

        this.update()
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

        const geometry = new THREE.SphereGeometry(50.0)
        const material = new THREE.MeshBasicMaterial(
            { color: 0x00ff00, wireframe: false }
        )
        const mesh = new THREE.Mesh(geometry, material)
        const object = new THREE.Object3D()
        object.add(mesh)
        object.position.set(500,0,0)
        this.scene.add(object)

    }

    private createCamera(): void {
        const aspectRatio: number = this.calculateAspectRatio()
        this.camera = new THREE.PerspectiveCamera(60, aspectRatio, 1, 1100)
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

    private createControls(): void {
        const canvas: HTMLCanvasElement = this.canvas()

        this.controls = new FlyControls(canvas)
        this.controls.init()

        if (this.onMovement) {
            this.subscription = this.controls.movement.subscribe(this.onMovement)
        }
    }

    private createPlayer(): void {
        this.player = new Player()
        this.player.init()

        if (this.scene && this.player.object) {
            this.scene.add(this.player.object)
        }
    }

    private createFollowCamera(): void {
        if (!this.camera || !this.player) {
            return
        }

        this.followCamera = new FollowCamera(this.camera, this.player)
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
            if (this.running) {
                this.update()
            }
        })

        // get the delta for this loop
        const delta: number = this.clock.getDelta()

        // update the controls
        if (this.controls) {
            this.controls.update(delta)
        }

        // update the player
        if (this.player) {
            this.player.update(delta)

            if (this.player.object) {
                this.stats.playerLocation.copy(this.player.object.position)
                this.stats.playerRotation.copy(this.player.object.rotation)
                this.stats.playerQuaternion.copy(this.player.object.quaternion)
            }
        }

        if (this.camera) {
            this.stats.cameraLocation.copy(this.camera.position)
            this.stats.cameraRotation.copy(this.camera.rotation)
        }

        // update the follow camera
        if (this.followCamera) {
            this.followCamera.update(delta)
        }

        this.render()
    }

    private render(): void {
        if (this.renderer && this.scene && this.camera) {
            this.renderer.render(this.scene, this.camera)
        }
    }

    sceneTransform(fn: (scene: THREE.Scene) => void): void {
        if (this.scene) {
            fn(this.scene)
        }
    }

    updatePlayer(position: THREE.Vector3, rotation: THREE.Vector3): void {
        if (this.player) {
            this.player.update0(position, rotation)
        }
    }

    ngOnDestroy(): void {
        // disable update loop
        this.running = false

        // unsubscribe
        if (this.subscription) {
            this.subscription.unsubscribe()
        }

        // dispose the player
        if (this.player) {
            this.player.dispose()
        }

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