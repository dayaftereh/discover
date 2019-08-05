import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { AdminServiceModule } from "src/app/services/api/admin/admin-service.module";
import { TableModule } from 'primeng/table';

@NgModule({
    imports: [
        // angular
        BrowserModule,
        // primeng
        TableModule,
        // custom
        AdminServiceModule
    ]
})
export class AdminStarSystemsComponentModule {

}