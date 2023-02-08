import { NgModule} from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app-component/app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import {MatSidenavModule} from '@angular/material/sidenav';
import { SideNavComponent } from './side-nav/side-nav/side-nav.component';
import { HomeComponent } from './home/home.component';
import { BodyComponent } from "./body/body.component";
import { CalendarComponent } from './calendar/calendar.component';
import { FullCalendarComponent } from './full-calendar/full-calendar.component';


@NgModule({
    declarations: [
        AppComponent,
        BodyComponent,
        SideNavComponent,
        HomeComponent,
        CalendarComponent,
        FullCalendarComponent,
    ],
    exports: [
        MatSidenavModule
    ],
    providers: [],
    bootstrap: [AppComponent],
    imports: [
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        MatSlideToggleModule,
        MatSidenavModule
    ]
})
export class AppModule { }
