import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { BehaviorSubject, lastValueFrom } from 'rxjs';
import { Router } from '@angular/router';

const httpOptions = {
  headers: new HttpHeaders({
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})
export class FriendService {

  errorSubject: any = new BehaviorSubject<any>(null);
  errorMessage: any = this.errorSubject.asObservable();

  constructor(
    private http: HttpClient,
    private router: Router,
  ) { }

  requestFriend(Reciever : string): any {
    lastValueFrom(this.http.post('http://localhost:3000/RequestFriend', { "requester" : sessionStorage.getItem('username'), "reciever" : Reciever })).then(async (res: any) => {
      if (res) {
        this.errorSubject.next(null);
      } else if (res.Message) {
        this.errorSubject.next(res.Message);
      }
    });
  }

  acceptFriend(Reciever : string): any {
    lastValueFrom(this.http.post('http://localhost:3000/AcceptFriend', { "requester" : sessionStorage.getItem('username'), "reciever" : Reciever })).then(async (res: any) => {
      if (res) {
        this.errorSubject.next(null);
      } else if (res.Message) {
        this.errorSubject.next(res.Message);
      }
    });
  }
}
