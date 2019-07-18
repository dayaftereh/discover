import { NgModule } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { BrowserModule } from "@angular/platform-browser";
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { LoginComponent } from "./login.component";

@NgModule({
    imports: [
        // angular
        FormsModule,
        BrowserModule,
        // primeng
        CardModule,
        ButtonModule,
        InputTextModule
    ],
    declarations: [
        LoginComponent
    ],
})
export class LoginComponentModule {

}