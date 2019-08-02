import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { FileUploadModule } from 'primeng/fileupload';
import { ThreeJSComponentModule } from "../game/threejs/threejs-component.module";
import { PlanetPreviewComponent } from "./planet-preview.component";

@NgModule({
    imports: [
        //angular
        BrowserModule,
        // primeng
        FileUploadModule,
        // custom
        ThreeJSComponentModule
    ],
    declarations: [
        PlanetPreviewComponent
    ],
    exports: [
        PlanetPreviewComponent
    ]
})
export class PlanetPreviewComponentModule {

}