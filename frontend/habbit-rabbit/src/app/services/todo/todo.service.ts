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

  addTask(task: string): any{
    lastValueFrom(this.http.post(this.url, {"user": sessionStorage.getItem('userID'), "description": task })).then(async (res: any) =>{
      if (res) {
     
        this.errorSubject.next(null);
     
      }
    });
  }

  editTask(task: string): any{
    lastValueFrom(this.http.put(this.url, { "user": sessionStorage.getItem('userID'), "description": task })).then(async (res: any) =>{
      if (res) {
     
        this.errorSubject.next(null);
     
      }
    });
  }

  /*deleteTask(task: string): any{
    lastValueFrom(this.http.delete(this.url, { "user": sessionStorage.getItem('userID'), "description": task })).then(async (res: any) =>{
      if (res) {
     
        this.errorSubject.next(null);
     
      }
    });
  }*/
}