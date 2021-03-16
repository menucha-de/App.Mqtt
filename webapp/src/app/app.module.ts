import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { MicaControlsModule } from '@menucha-de/controls';
import { MicaAppBaseModule, MicaAppComponentsModule } from '@menucha-de/shared';
import { AccControlComponent } from './acc-control/acc-control.component';
import { HttpClientModule } from '@angular/common/http';
import { MainComponent } from './main/main.component';
import { SecurityComponent } from './security/security.component';
import { FormsModule } from '@angular/forms';


@NgModule({
  declarations: [
    AppComponent,
    MainComponent,
    AccControlComponent,
    SecurityComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    MicaControlsModule,
    MicaControlsModule,
    MicaAppBaseModule,
    MicaAppComponentsModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
