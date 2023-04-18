import { Component } from '@angular/core';
import { SideNavComponent } from '../side-nav/side-nav/side-nav.component';
import { Router } from '@angular/router';
import { ToDoService } from '../services/todo/todo.service';

export interface Item {
  description: string;
  done: boolean;
}

export interface SideNavToggle{
  screenWidth: number;
  collapsed: boolean;
}

@Component({
  selector: 'app-todolist',
  templateUrl: './todolist.component.html',
  styleUrls: ['./todolist.component.css']
})

export class TodolistComponent {
  isSideNavCollapsed = false;
  screenWidth=0;
  error: any = null;
  onToggleSideNav(data: SideNavToggle): void{
    this.screenWidth = data.screenWidth;
    this.isSideNavCollapsed = data.collapsed;
  }

  constructor(
    private todoService: ToDoService, private router: Router
  ) { }

  ngOnInit(): void {
    this.todoService
      .errorSubject
      .subscribe((errorMessage: any) => {
        this.error = errorMessage;
      });
  }

  filter: 'all' | 'active' | 'done' = 'all';

  allItems = [
    { description: 'Record SWE video', done: false }
  ];

  get items() {
    if (this.filter === 'all') {
      return this.allItems;
    }
    return this.allItems.filter((item) => this.filter === 'done' ? item.done : !item.done);
  }

  completeItem(itemObj: Item){
    itemObj.done=true;
    this.todoService
      .editTask(itemObj.description);
    
  }

  addItem(description: string) {
    this.allItems.unshift({
      description,
      done: false
    });
    this.todoService
      .addTask(description);
  }

  deleteItem(itemObj: Item){
    this.allItems.splice(this.allItems.indexOf(itemObj), 1);
    this.todoService
      .deleteTask(itemObj.description);
  }
}
