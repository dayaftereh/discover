import { NgModule } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { BrowserModule } from "@angular/platform-browser";
import { CheckboxModule } from 'primeng/checkbox';
import { FileUploadModule } from 'primeng/fileupload';
import { ThreeJSComponentModule } from "../game/threejs/threejs-component.module";
import { PlanetPreviewComponent } from "./planet-preview.component";

@NgModule({
    imports: [
        //angular
        FormsModule,
        BrowserModule,
        // primeng
        CheckboxModule,
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