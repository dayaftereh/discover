import { NgModule } from "@angular/core";
import { CommonServiceModule } from "../api/common/common-service.module";
import { RouterModule } from "@angular/router";
import { AuthenticationGuardService } from "./authentication-guard.service";
import { InGameGuardService } from "./in-game-guard.service";

@NgModule({
    imports: [
        //angular
        RouterModule,
        // custom
        CommonServiceModule,
    ],
    providers: [
        InGameGuardService,
        AuthenticationGuardService
    ]
})
export class GuardServicesModule {

}