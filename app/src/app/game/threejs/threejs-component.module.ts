import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { ThreeJSComponent } from "./threejs.component";

@NgModule({
    imports: [
        // angular
        BrowserModule
    ],
    declarations: [
        ThreeJSComponent
    ],
    exports: [
        ThreeJSComponent
    ]
})
export class ThreeJSComponentModule {

}