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
  errorSubject: any = new BehaviorSubject<any>(null);
  errorMessage: any = this.errorSubject.asObservable();

  constructor(
    private http: HttpClient,
    private router: Router,
  ) { }

  addTask(task: string): any{
    lastValueFrom(this.http.post('http://localhost:3000/EditToDo', {"user": sessionStorage.getItem('userID'), "description": task })).then(async (res: any) =>{
      if (res) {
     
        this.errorSubject.next(null);
     
      }
    });
  }

  editTask(task: string): any{
    lastValueFrom(this.http.put('http://localhost:3000/EditToDo', { "user": sessionStorage.getItem('userID'), "description": task })).then(async (res: any) =>{
      if (res) {
        this.errorSubject.next(null);
      } else if (res.Message) {
        this.errorSubject.next(res.Message);
      }
    });
  }

  /*deleteTask(task: string): any{
    lastValueFrom(this.http.delete('http://localhost:3000/EditToDo', { "user": sessionStorage.getItem('userID'), "description": task })).then(async (res: any) =>{
      if (res) {
     
        this.errorSubject.next(null);
     
      }
    });
  }*/
}