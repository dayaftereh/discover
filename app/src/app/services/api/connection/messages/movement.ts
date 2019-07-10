import { Message } from "./message";

export interface Movement extends Message {
    move: {
        x: number,
        y: number,
        z: number
    },
    rotation: {
        x: number,
        y: number,
        z: number
    }
}