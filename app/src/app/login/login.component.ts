import { Component, OnInit } from "@angular/core";
import { CommonService } from "../services/api/common/common.service";
import { Status } from "../services/api/common/status";
import { Router } from "@angular/router";

@Component({
    templateUrl: './login.component.html'
})
export class LoginComponent implements OnInit {

    username: string | undefined

    constructor(private readonly commonService: CommonService,
        private readonly router: Router) { }

    async ngOnInit(): Promise<void> {
        try {
            const status: Status = await this.commonService.status()
            if (status && status.name) {
                await this.gotoGame()
            }
        } catch (e) {
            console.log("not login")
        }
    }


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