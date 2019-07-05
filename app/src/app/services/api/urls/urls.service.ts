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

    websocketUrl(): string {
        return this.apiJoin("/ws")
    }

}