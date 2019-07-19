import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { DropdownModule } from 'primeng/dropdown';
import { MenubarModule } from 'primeng/menubar';
import { CommonServiceModule } from "../services/api/common/common-service.module";
import { NavbarComponent } from "./navbar.component";
import { ButtonModule } from "primeng/button";

@NgModule({
    imports: [
        // angulat
        BrowserModule,
        // primeng
        MenubarModule,  
        ButtonModule,     
        //custom
        CommonServiceModule
    ],
    declarations: [
        NavbarComponent
    ],
    exports: [
        NavbarComponent
    ]
})
export class NavbarComponentModule { }