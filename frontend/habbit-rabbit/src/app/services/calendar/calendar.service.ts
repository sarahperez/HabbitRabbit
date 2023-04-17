import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { BehaviorSubject, lastValueFrom} from 'rxjs';
import { Router } from '@angular/router';
import { CalendarComponent } from 'src/app/calendar/calendar.component';
//import { getEventNum } from 'src/app/calendar/calendar.component';

const httpOptions = {
  headers: new HttpHeaders({
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})
export class CalendarService {
  url: any = 'http://localhost:3000/calendar';
  errorSubject: any = new BehaviorSubject<any>(null);
  errorMessage: any = this.errorSubject.asObservable();

  constructor(
    private http: HttpClient,
    private router: Router,
  ) { }

  addEvent(EventID: string, Start: string, End: string, Title: string): any {
    lastValueFrom(this.http.post(this.url, { "user" : sessionStorage.getItem('userID'), "eventID" : EventID, "startStr" : Start, "endStr": End, "title" : Title })).then(async (res: any) => {
      if (res) {
        sessionStorage.setItem('jwt', res.jwt);
        this.errorSubject.next(null);
      } 
    });
  }
/*
  deleteEvent(EventID: string, Start: string, End: string, Title: string): any {
    lastValueFrom(this.http.delete(this.url, { "user" : sessionStorage.getItem('userID'), "eventID" : EventID, "startStr" : Start, "endStr": End, "title" : Title })).then(async (res: any) => {
      if (res&&res.jwt) {
        sessionStorage.setItem('jwt', res.jwt);
        this.errorSubject.next(null);
      } else if (res.Message) {
        this.errorSubject.next(res.Message);
      }
    });
  }
*/

  loadEvents(): any {
    lastValueFrom(this.http.post(this.url, { "user": sessionStorage.getItem('userID')})).then(async (res: any) => {
      if (res) {
        sessionStorage.setItem('events', res.data['items']);
      }
    })
  }

}
