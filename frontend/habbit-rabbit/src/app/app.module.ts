import { NgModule} from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app-component/app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { SideNavComponent } from './side-nav/side-nav/side-nav.component';
import { HomeComponent } from './home/home.component';
import { BodyComponent } from "./body/body.component";
import { CalendarComponent } from './calendar/calendar.component';
import { FullCalendarModule } from '@fullcalendar/angular';
import { FullCalendarComponent } from './full-calendar/full-calendar.component';
import { LoginComponent } from './login/login.component';
import { TodolistComponent } from './todolist/todolist.component';
import { FormsModule } from '@angular/forms';
import { FriendsComponent } from './friends/friends.component';

@NgModule({
    declarations: [
        AppComponent,
        BodyComponent,
        SideNavComponent,
        HomeComponent,
        CalendarComponent,
        FullCalendarComponent,
        LoginComponent,
        TodolistComponent,
        FriendsComponent
    ],
    exports: [
        FullCalendarModule
    ],
    providers: [],
    bootstrap: [AppComponent],
    imports: [
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        FullCalendarModule,
        HttpClientModule,
        FormsModule
    ]
})
export class AppModule { }
