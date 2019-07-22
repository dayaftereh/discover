import { Message } from "./message";

export interface Pong extends Message {
    clientSendTime: number
    serverReceiveTime: number
    serverSendTime: number
}