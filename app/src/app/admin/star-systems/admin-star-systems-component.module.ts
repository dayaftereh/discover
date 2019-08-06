import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { RouterModule } from "@angular/router";
import { ButtonModule } from 'primeng/button';
import { PanelModule } from 'primeng/panel';
import { TableModule } from 'primeng/table';
import { AdminServiceModule } from "src/app/services/api/admin/admin-service.module";
import { AdminStarSystemComponent } from "./admin-star-system.component";
import { AdminStarSystemsComponent } from "./admin-star-systems.component";
import { AdminPlanetComponent } from "./admin-planet.component";
import { CheckboxModule } from "primeng/checkbox";
import { FormsModule } from "@angular/forms";

@NgModule({
    imports: [
        // angular
        FormsModule,
        RouterModule,
        CommonModule,
        BrowserModule,
        // primeng
        TableModule,
        PanelModule,
        ButtonModule,
        CheckboxModule,
        // custom
        AdminServiceModule
    ],
    declarations: [
        AdminPlanetComponent,
        AdminStarSystemComponent,
        AdminStarSystemsComponent
    ]
})
export class AdminStarSystemsComponentModule {

}