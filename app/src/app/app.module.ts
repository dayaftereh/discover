import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { LayoutModule } from './layout/layout.module';
import { AppRouting } from './app.routing';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    // angular
    BrowserModule,
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
