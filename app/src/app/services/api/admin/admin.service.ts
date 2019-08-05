import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { URLSService } from "../urls/urls.service";
import { StarSystem } from "./star-system";

@Injectable()
export class AdminService {

    constructor(
        private readonly http: HttpClient,
        private readonly urlsService: URLSService,
    ) { }


    async allStarSystems(): Promise<string[]> {
        const url: string = this.urlsService.apiJoin('admin', 'star-systems')
        const starSystems: string[] = await this.http.get<string[]>(url).toPromise()
        return starSystems
    }

    async starSystem(name: string): Promise<StarSystem> {
        const url: string = this.urlsService.apiJoin('admin', 'star-system', name)
        const starSystem: StarSystem = await this.http.get<StarSystem>(url).toPromise()
        return starSystem
    }

}