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

  requestFriend(Receiver : string): any {
    lastValueFrom(this.http.post('http://localhost:3000/RequestFriend', { "requester" : sessionStorage.getItem('loggedUsername'), "reciever" : Receiver})).then(async (res: any) => {
      if (res) {
        this.errorSubject.next(null);
      } else if (res.Message) {
        this.errorSubject.next(res.Message);
      }
    });
  }

  acceptFriend(Sender : string): any {
    lastValueFrom(this.http.post('http://localhost:3000/AcceptFriend', { "requester" : Sender, "reciever" : sessionStorage.getItem('loggedUsername')})).then(async (res: any) => {
      if (res) {
        this.errorSubject.next(null);
      } else if (res.Message) {
        this.errorSubject.next(res.Message);
      }
    });
  }

  blockFriend(Sender : string): any {
    lastValueFrom(this.http.post('http://localhost:3000/AcceptFriend', { "requester" : Sender, "reciever" : sessionStorage.getItem('loggedUsername')})).then(async (res: any) => {
      if (res) {
        this.errorSubject.next(null);
      } else if (res.Message) {
        this.errorSubject.next(res.Message);
      }
    });
  }

  getFriendStatus(): any {
    lastValueFrom(this.http.post('http://localhost:3000/FriendStatus', { "user": +sessionStorage.getItem('userID')!})).then(async (res: any) => {
      if (res) {
        sessionStorage.setItem('pendingRequests', res.data['Requests from']);
        sessionStorage.setItem('blocked', res.data['Blocked Users']);
        sessionStorage.setItem('friends', res.data['Friends']);
        this.errorSubject.next(null);
      } else if (res.Message) {
        this.errorSubject.next(res.Message);
      }
    })
  }
}
