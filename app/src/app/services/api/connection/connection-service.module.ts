import { NgModule } from "@angular/core";
import { ConnectionService } from "./connection.service";
import { URLSServiceModule } from "../urls/urls-service.module";

@NgModule({
    imports: [
        //customs
        URLSServiceModule
    ],
    providers: [
        ConnectionService
    ]
})
export class ConnectionServiceModule {

}