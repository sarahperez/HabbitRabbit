import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { BehaviorSubject, lastValueFrom} from 'rxjs';
import { Router } from '@angular/router';

const httpOptions = {
  headers: new HttpHeaders({
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})
export class ToDoService {
  url: any = 'http://localhost:3000/todo';
  errorSubject: any = new BehaviorSubject<any>(null);
  errorMessage: any = this.errorSubject.asObservable();

  constructor(
    private http: HttpClient,
    private router: Router,
  ) { }

  


  login(Username: string, Password: string): any {
    lastValueFrom(this.http.post(this.url, { "username": Username, "password": Password })).then(async (res: any) => {
      if (res&&res.jwt) {
        sessionStorage.setItem('loggedUser', res.data['Name']);
        sessionStorage.setItem('jwt', res.jwt);
        this.errorSubject.next(null);
        this.router.navigateByUrl('home');
      } else if (res.Message) {
        this.errorSubject.next(res.Message);
      }
    });
  }
}