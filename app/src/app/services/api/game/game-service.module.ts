import { NgModule } from "@angular/core";
import { GameService } from "./game.service";
import { URLSServiceModule } from "../urls/urls-service.module";
import { HttpClientModule } from "@angular/common/http";
import { BrowserModule } from "@angular/platform-browser";

@NgModule({
    imports: [
        // angular
        BrowserModule,
        HttpClientModule,
        //custom
        URLSServiceModule
    ],
    providers: [
        GameService
    ]
})
export class GameServiceModule {

}