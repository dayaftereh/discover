import { Injectable } from "@angular/core";
import { Status } from "./status";
import { HttpClient } from "@angular/common/http";

@Injectable()
export class CommonService {

    constructor(
        private readonly http: HttpClient
    ) { }

    async status(): Promise<Status> {
        const status: Status = await this.http.get<Status>('common/status').toPromise()
        return status
    }

    async login(name: string): Promise<Status> {
        const status: Status = await this.http.post<Status>('common/status', {
            name
        }).toPromise()
        return status
    }

    async logout(): Promise<void> {
        await this.http.post('common/logout', {})
    }

}