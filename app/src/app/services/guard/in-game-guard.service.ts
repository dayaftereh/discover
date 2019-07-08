import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree, Router } from "@angular/router";
import { Injectable } from "@angular/core";
import { CommonService } from "../api/common/common.service";

@Injectable()
export class InGameGuardService implements CanActivate {

    constructor(private readonly router: Router,
        private readonly commonService: CommonService) {

    }

    async canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Promise<UrlTree | boolean> {
        // check if user is Authenticated
        const isAuth: boolean = await this.commonService.isAuthenticated()
        if (!isAuth) {
            return true
        }

        // goto game, beacuse user is already ingame
        return this.router.createUrlTree(["/game"])
    }



}