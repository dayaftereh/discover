import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppComponent } from './app.component';
import { AppRouting } from './app.routing';
import { LayoutModule } from './layout/layout.module';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    // angular
    BrowserModule,
    BrowserAnimationsModule,
    // Routing
    AppRouting,
    // custom
    LayoutModule,
  ],
  bootstrap: [
    AppComponent
  ]
})
export class AppModule { }
