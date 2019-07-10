import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { URLSService } from "../urls/urls.service";

@Injectable()
export class GameService {

    constructor(
        private readonly http: HttpClient,
        private readonly urlsService: URLSService
    ) { }

    async ready(): Promise<void> {
        const url: string = this.urlsService.apiJoin('game', 'ready')
        await this.http.post(url, {}).toPromise()
    }

}