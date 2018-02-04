import { AppModule } from './app.module';
import {
  BrowserModule,
  BrowserTransferStateModule
} from '@angular/platform-browser';
import { NgModule } from '@angular/core';
// Installed separately
import { ServerModule } from '@angular/platform-server';

import { AppComponent } from './app.component';

@NgModule({
  declarations: [],
  imports: [
    // Make sure the string matches
    BrowserModule.withServerTransition({
      appId: 'my-app-id'
    }),
    ServerModule,
    AppModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class ServerAppModule {}
