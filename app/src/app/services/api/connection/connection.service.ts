import { Subject, Subscription } from "rxjs";
import { URLSService } from "../urls/urls.service";
import { Message } from "./messages/message";

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
                    const obj: any = JSON.parse(event.data)
                    if (obj){
                        this.subject.next(obj)
                    }
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

    onMessage(fn: (message: Message) => void): Subscription | undefined {
        if (this.subject) {
            return this.subject.subscribe(fn)
        }
        return undefined
    }

    send(message: Message): void {
        const data: string = JSON.stringify(message)
        if (this.websocket && this.websocket.readyState === WebSocket.OPEN) {
            this.websocket.send(data)
        }
    }

}