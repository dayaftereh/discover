import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { InputTextModule } from 'primeng/inputtext';
import { FormsModule } from "@angular/forms";
import { ButtonModule } from 'primeng/button';
import { LoginComponent } from "./login.component";

@NgModule({
    imports: [
        // angular
        FormsModule,
        BrowserModule,
        // primeng
        ButtonModule,
        InputTextModule
    ],
    declarations:[
        LoginComponent
    ],
})
export class LoginComponentModule {

}