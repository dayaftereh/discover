import { Message } from "./message";

export interface Ping extends Message {
    clientTime: number
}