import { Message } from "./message";

export interface Vec3 {
    x: number
    y: number
    z: number
}

export interface GameObject {
    radius?: number
    position?: Vec3
    rotation?: Vec3
    removeable?: boolean
}

export interface Update extends Message {
    tick: number
    time: number
    player: GameObject
    objects: { [key: number]: GameObject }
}