import { Message } from "./message";

export interface Vec3 {
    x: number
    y: number
    z: number
}

export interface GameObject {
    type?:string
    radius?: number
    position?: Vec3
    rotation?: Vec3
    removeable?: boolean
    color?: number
}

export interface PlayerUpdate {
    gameObjectId: number
}

export interface WorldUpdate extends Message {
    tick: number
    time: number
    player: PlayerUpdate
    objects: { [key: number]: GameObject }
}