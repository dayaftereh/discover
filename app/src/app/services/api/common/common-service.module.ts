import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { HttpClientModule } from '@angular/common/http';
import { CommonService } from "./common.service";
import { URLSServiceModule } from "../urls/urls-service.module";

@NgModule({
    imports: [
        // angular
        BrowserModule,
        HttpClientModule,
        //custom
        URLSServiceModule
    ],
    providers: [
        CommonService
    ]
})
export class CommonServiceModule {

}