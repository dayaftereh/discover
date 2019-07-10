import { Injectable } from "@angular/core";
import { Status } from "./status";
import { HttpClient } from "@angular/common/http";
import { URLSService } from "../urls/urls.service";

@Injectable()
export class CommonService {

    constructor(
        private readonly http: HttpClient,
        private readonly urlsService: URLSService
    ) { }

    async isAuthenticated(): Promise<boolean> {
        // get the status
        const status: Status = await this.status()
        // check if authenticated
        return status.authenticated
    }

    async status(): Promise<Status> {
        const url: string = this.urlsService.apiJoin('common', 'status')
        const status: Status = await this.http.get<Status>(url).toPromise()
        return status
    }

    async login(name: string): Promise<Status> {
        const url: string = this.urlsService.apiJoin('common', 'login')
        const status: Status = await this.http.post<Status>(url, {
            name
        }).toPromise()
        return status
    }

    async logout(): Promise<void> {
        const url: string = this.urlsService.apiJoin('common', 'logout')
        await this.http.post(url, {}).toPromise()
    }

}