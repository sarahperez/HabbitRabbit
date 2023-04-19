import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SideNavComponent } from '../side-nav/side-nav/side-nav.component';
import { TodolistComponent } from './todolist.component';
import { ToDoService } from '../services/todo/todo.service';
import { Router } from '@angular/router';

describe('TodolistComponent', () => {
  let component: TodolistComponent;
  let fixture: ComponentFixture<TodolistComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TodolistComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TodolistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

describe('Todo List Functionality', () => {
  let component: TodolistComponent;
  let fixture: ComponentFixture<TodolistComponent>;

  it('should add a list item', () => {
    component.addItem('New item');
    expect(component.allItems.length).toBeGreaterThan(3);
  });

  it('should delete a list item', () => {
    //component.deleteItem(new itemObj);
    expect(component.allItems.length).toBeGreaterThan(3);
  });
});
