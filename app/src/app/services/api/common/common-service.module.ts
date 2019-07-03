import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { HttpClientModule } from '@angular/common/http';
import { CommonService } from "./common.service";

@NgModule({
    imports: [
        // angular
        BrowserModule,
        HttpClientModule
    ],
    providers: [
        CommonService
    ]
})
export class CommonServiceModule {

}