import { Component, OnInit } from "@angular/core";
import { CommonService } from "../services/api/common/common.service";
import { Status } from "../services/api/common/status";
import { Router } from "@angular/router";

@Component({
    templateUrl: './login.component.html',
    styleUrls:[
        './login.component.scss'
    ]
})
export class LoginComponent {

    username: string | undefined

    constructor(private readonly commonService: CommonService,
        private readonly router: Router) { }

    async login(): Promise<void> {
        if (!this.username) {
            return
        }
        // login the user
        const status: Status = await this.commonService.login(this.username)

        await this.gotoGame()
    }

    async gotoGame(): Promise<void> {
        await this.router.navigate(["game"])
    }

}