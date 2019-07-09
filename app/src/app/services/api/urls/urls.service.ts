import { Injectable } from "@angular/core";

@Injectable()
export class URLSService {

    constructor() { }

    apiJoin(...args: string[]): string {
        return this.join('api', ...args)
    }

    join(...args: string[]): string {
        return args.map((part: string, i: number) => {
            if (i === 0) {
                return part.trim().replace(/[\/]*$/g, '')
            } else {
                return part.trim().replace(/(^[\/]*|[\/]*$)/g, '')
            }
        }).filter((x: string) => x.length > 0).join('/')
    }

    apiWebsocketUrl(): string {
        return this.websocketUrl("/api/ws")
    }

    private websocketUrl(path: string): string {
        const url: URL = new URL(path, window.location.href)
        url.protocol = url.protocol.replace('http', 'ws')
        return url.href
    }

}