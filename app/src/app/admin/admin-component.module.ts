import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { AdminComponent } from "./admin.component";
import { AdminStarSystemsComponentModule } from "./star-systems/admin-star-systems-component.module";

@NgModule({
    imports: [
        // angular
        BrowserModule,
        //custom
        AdminStarSystemsComponentModule
    ],
    declarations: [
        AdminComponent
    ]
})
export class AdminComponentModule {

}