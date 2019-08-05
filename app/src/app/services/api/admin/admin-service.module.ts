import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { HttpClientModule } from "@angular/common/http";
import { URLSServiceModule } from "../urls/urls-service.module";
import { AdminService } from "./admin.service";

@NgModule({
    imports: [
        // angular
        BrowserModule,
        HttpClientModule,
        //custom
        URLSServiceModule
    ],
    providers: [
        AdminService
    ]
})
export class AdminServiceModule{

}