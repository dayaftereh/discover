import * as THREE from 'three';
import { ControlsState } from "./controls-state";
import { Subject } from 'rxjs';
import { Movement } from '../../services/api/connection/messages/movement';
import { MessageType } from 'src/app/services/api/connection/messages/message-type';
import { GameComponent } from './game-component';
import { EventEmitter } from '@angular/core';
import { GameInputEvent } from './game-input-event';
import { ConnectionService } from 'src/app/services/api/connection/connection.service';

export class GameInput implements GameComponent {

    private move: THREE.Vector3
    private rotation: THREE.Vector3

    private controlsState: ControlsState

    constructor(private readonly element: HTMLElement,
        private readonly connectionService: ConnectionService) {

        this.move = new THREE.Vector3()
        this.rotation = new THREE.Vector3()

        this.controlsState = {
            thruster: 0,
            strafe: 0,
            up: 0,

            pitch: 0,
            yaw: 0,
            roll: 0,
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

        this.element.addEventListener('mouseout', () => {
            this.onMouseOut()
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

        this.controlsState.yaw = -((x - dimension.offset.x) - halfWidth) / halfWidth
        this.controlsState.pitch = ((y - dimension.offset.y) - halfHeight) / halfHeight

        this.updateRotationVector()
    }

    private onMouseOut(): void {
        this.rotation.x = 0.0
        this.rotation.y = 0.0
        this.rotation.z = 0.0
    }

    private onKeyDown(event: KeyboardEvent): void {
        if (this.handleKeyEvent(event.key, 1.0)) {
            event.preventDefault()
        }
    }

    private onKeyUp(event: KeyboardEvent): void {
        if (this.handleKeyEvent(event.key, 0.0)) {
            event.preventDefault()
        }
    }

    private handleKeyEvent(key: string, state: number): boolean {
        let consumed: boolean = true;
        const keyName: string = key.toLowerCase()
        switch (keyName) {
            case 'w':
                this.controlsState.thruster = state; break;
            case 's':
                this.controlsState.thruster = -state; break;

            case 'a':
                this.controlsState.strafe = state; break;
            case 'd':
                this.controlsState.strafe = -state; break;

            case 'r':
                this.controlsState.up = state; break;
            case 'f':
                this.controlsState.up = -state; break;

            case 'arrowup':
                this.controlsState.pitch = -state; break;
            case 'arrowdown':
                this.controlsState.pitch = state; break;

            case 'arrowleft':
                this.controlsState.yaw = state; break;
            case 'arrowright':
                this.controlsState.yaw = -state; break;

            case 'q':
                this.controlsState.roll = -state; break;
            case 'e':
                this.controlsState.roll = state; break;
            default:
                consumed = false; break;
        }

        if (consumed) {
            this.updateMovementVector()
            this.updateRotationVector()
        }

        return consumed
    }

    private updateMovementVector(): void {
        this.move.x = this.controlsState.strafe
        this.move.y = this.controlsState.up
        this.move.z = this.controlsState.thruster

    }

    private updateRotationVector(): void {
        this.rotation.x = this.controlsState.pitch
        this.rotation.y = this.controlsState.yaw
        this.rotation.z = this.controlsState.roll
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
        this.connectionService.send(movement)
    }

    dispose(): void {

    }

}