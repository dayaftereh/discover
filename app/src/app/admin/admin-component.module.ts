import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { AdminStarSystemsComponentModule } from "./star-systems/admin-star-systems-component.module";
import { TabMenuModule } from 'primeng/tabmenu';

@NgModule({
    imports: [
        // angular
        BrowserModule,
        //primeng
        TabMenuModule,
        //custom
        AdminStarSystemsComponentModule
    ]
})
export class AdminComponentModule {

}