import * as THREE from 'three';
import { ControlsState } from "./controls-state";
import { Subject } from 'rxjs';
import { Movement } from '../../services/api/connection/messages/movement';
import { MessageType } from 'src/app/services/api/connection/messages/message-type';

export class FlyControls {

    movement: Subject<Movement>

    private move: THREE.Vector3
    private rotation: THREE.Vector3

    private controlsState: ControlsState

    constructor(private readonly element: HTMLElement) {
        this.movement = new Subject()

        this.move = new THREE.Vector3()
        this.rotation = new THREE.Vector3()

        this.controlsState = {
            forward: 0.0,
            backward: 0.0,
            left: 0.0,
            right: 0.0,
            up: 0.0,
            down: 0.0,
            pitchUp: 0.0,
            pitchDown: 0.0,
            yawLeft: 0.0,
            yawRight: 0.0,
            rollLeft: 0.0,
            rollRight: 0.0
        } as ControlsState
    }

    init(): void {
        window.addEventListener('keyup', (event: KeyboardEvent) => {
            this.onKeyUp(event)
        })

        window.addEventListener('keydown', (event: KeyboardEvent) => {
            this.onKeyDown(event)
        })

        this.element.addEventListener('mousemove', (event: MouseEvent) => {
            this.onMouseMove(event)
        })

        this.element.addEventListener('contextmenu', (event: MouseEvent) => {
            event.preventDefault()
        })
    }

    private onMouseMove(event: MouseEvent): void {
        const x: number = event.pageX
        const y: number = event.pageY

        const dimension = this.getContainerDimensions()

        const halfWidth: number = dimension.size.width / 2.0
        const halfHeight: number = dimension.size.height / 2.0

        this.controlsState.yawLeft = -((x - dimension.offset.x) - halfWidth) / halfWidth
        this.controlsState.pitchDown = ((y - dimension.offset.y) - halfHeight) / halfHeight

        this.updateRotationVector()
    }

    private onKeyDown(event: KeyboardEvent): void {
        this.handleKeyEvent(event.key, 1.0)
    }

    private onKeyUp(event: KeyboardEvent): void {
        this.handleKeyEvent(event.key, 0.0)
    }

    private handleKeyEvent(key: string, state: number): void {
        console.log("Key:", key)
        const keyName: string = key.toLowerCase()
        switch (keyName) {
            case 'w':
                this.controlsState.forward = state; break;
            case 's':
                this.controlsState.backward = state; break;

            case 'a':
                this.controlsState.left = state; break;
            case 'd':
                this.controlsState.right = state; break;

            case 'r':
                this.controlsState.up = state; break;
            case 'f':
                this.controlsState.down = state; break;

            case 'arrowup':
                this.controlsState.pitchUp = state; break;
            case 'arrowdown':
                this.controlsState.pitchDown = state; break;

            case 'arrowleft':
                this.controlsState.yawLeft = state; break;
            case 'arrowright':
                this.controlsState.yawRight = state; break;

            case 'q':
                this.controlsState.rollLeft = state; break;
            case 'e':
                this.controlsState.rollRight = state; break;
        }

        this.updateMovementVector()
        this.updateRotationVector()
    }

    private updateMovementVector(): void {
        this.move.x = (-this.controlsState.left + this.controlsState.right)
        this.move.y = (-this.controlsState.down + this.controlsState.up)
        this.move.z = (-this.controlsState.forward + this.controlsState.backward)

    }

    private updateRotationVector(): void {
        this.rotation.x = (-this.controlsState.pitchDown + this.controlsState.pitchUp)
        this.rotation.y = (-this.controlsState.yawRight + this.controlsState.yawLeft)
        this.rotation.z = (-this.controlsState.rollRight + this.controlsState.rollLeft)
    }

    private getContainerDimensions(): {
        size: {
            width: number,
            height: number
        },
        offset: {
            x: number,
            y: number
        }
    } {
        if (!this.element) {
            return {
                size: {
                    width: window.innerWidth,
                    height: window.innerHeight
                },
                offset: {
                    x: 0.0,
                    y: 0.0
                }
            }
        }

        return {
            size: {
                width: this.element.offsetWidth,
                height: this.element.offsetHeight
            },
            offset: {
                x: this.element.offsetLeft,
                y: this.element.offsetTop
            }
        }
    }

    update(delta: number): void {
        // create the movement
        const movement: Movement = {
            type: MessageType.MOVE,
            move: this.move,
            rotation: this.rotation
        }

        // notify about new movement
        this.movement.next(movement)
    }

}