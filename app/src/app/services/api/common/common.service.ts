import { Injectable } from "@angular/core";
import { Status } from "./status";
import { HttpClient } from "@angular/common/http";
import { URLSService } from "../urls/urls.service";
import { StatusService } from "./status.service";

@Injectable()
export class CommonService {

    constructor(
        private readonly http: HttpClient,
        private readonly urlsService: URLSService,
        private readonly statusService: StatusService,
    ) { }

    async isAuthenticated(): Promise<boolean> {
        // get the status
        const status: Status = await this.status()

        // fire the new status
        this.statusService.update(status)

        // check if authenticated
        return status.authenticated
    }

    async status(): Promise<Status> {
        const url: string = this.urlsService.apiJoin('common', 'status')
        const status: Status = await this.http.get<Status>(url).toPromise()

        // fire the new status
        this.statusService.update(status)

        return status
    }

    async login(name: string): Promise<Status> {
        const url: string = this.urlsService.apiJoin('common', 'login')
        const status: Status = await this.http.post<Status>(url, {
            name
        }).toPromise()

        // fire the new status
        this.statusService.update(status)
        
        return status
    }

    async logout(): Promise<void> {
        const url: string = this.urlsService.apiJoin('common', 'logout')
        await this.http.post(url, {}).toPromise()
    }

}