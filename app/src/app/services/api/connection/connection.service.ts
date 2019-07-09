import { Subject } from "rxjs";
import { Message } from "./message";
import { URLSService } from "../urls/urls.service";

export class ConnectionService {

    private websocket: WebSocket | undefined
    private subject: Subject<Message> | undefined

    constructor(private readonly urlsService: URLSService) {

    }

    isConnected(): boolean {
        return !!(this.websocket) && this.websocket.readyState == WebSocket.OPEN
    }

    connect(): Promise<void> {
        let completed: boolean = false

        // return a promise to wait for websocket open
        return new Promise((resolve, reject) => {
            this.subject = new Subject<Message>();

            // get the websocket url
            const url: string = this.urlsService.apiWebsocketUrl()
            // create the websocket
            this.websocket = new WebSocket(url)

            this.websocket!.addEventListener('open', () => {
                if (!completed) {
                    completed = true
                    resolve()
                }
            })

            this.websocket!.addEventListener('message', (event: MessageEvent) => {
                if (event.data && this.subject) {
                    this.subject.next(event.data)
                }
            })

            this.websocket!.addEventListener('error', (e) => {
                if (!completed) {
                    completed = true
                    return reject(e)
                }

                if (e && this.subject) {
                    this.subject.error(e)
                }
            })

            this.websocket!.addEventListener('close', () => {
                if (this.subject) {
                    this.subject.complete()
                }

                this.subject = undefined
                this.websocket = undefined
            })
        })
    }

    close(): void {
        if (this.websocket) {
            this.websocket.close()
        }
    }

    onMessage(fn: (message: Message) => void) {
        if (this.subject) {
            this.subject.subscribe(fn)
        }
    }

    send(message: Message): void {
        const data: string = JSON.stringify(message)
        if (this.websocket && this.websocket.readyState === WebSocket.OPEN) {
            this.websocket.send(data)
        }
    }

}