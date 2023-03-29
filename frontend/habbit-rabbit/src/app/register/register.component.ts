import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { RegisterService } from './../services/register/register.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  username: any;
  password: any;
  name: any;
  email: any;

  constructor(
    private registerService: RegisterService, private router: Router
  ) { }

  onKey(event: any, type: string) {
    if (type === 'name') {
      this.name = event.target.value;
    } else if (type === 'email') {
      this.email = event.target.value;
    } else if (type === 'username') {
      this.username = event.target.value;
    } else if (type === 'password') {
      this.password = event.target.value;
    }
  }

  onSubmit() {
    this.registerService
        .register(this.name, this.email, this.username, this.password);
  }
  
}
