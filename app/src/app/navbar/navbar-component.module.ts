import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { MenubarModule } from 'primeng/menubar';
import { NavbarComponent } from "./navbar.component";
import { DropdownModule } from 'primeng/dropdown';

@NgModule({
    imports: [
        // angulat
        BrowserModule,
        // primeng
        MenubarModule,
        DropdownModule,
    ],
    declarations: [
        NavbarComponent
    ],
    exports: [
        NavbarComponent
    ]
})
export class NavbarComponentModule { }