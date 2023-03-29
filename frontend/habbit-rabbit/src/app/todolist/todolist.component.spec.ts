import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SideNavComponent } from '../side-nav/side-nav/side-nav.component';
import { TodolistComponent } from './todolist.component';

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

describe('Todo Liast Functionality', () => {
  const component = new TodolistComponent();
  let fixture: ComponentFixture<TodolistComponent>;

  it('should add a list item', () => {
    component.addItem('New item');
    expect(component.allItems.length).toBeGreaterThan(1);
  });
});
