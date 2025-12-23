import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { NgIf } from '@angular/common';
import { Router, RouterLink } from '@angular/router';

import { TodoService } from '../services/todo.service';


@Component({
  selector: 'app-login',
  standalone: true,
  imports: [FormsModule, RouterLink, NgIf],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  email = '';
  password = '';
  response: any;
  errorMessage: string = '';

  constructor(private todoService: TodoService,  private router: Router) {}

  onLogin() {
    this.errorMessage = '';

    this.todoService.login({
      email: this.email,
      password: this.password
    }).subscribe({
      next: res => {
        this.response = res; 
        localStorage.setItem('token', res.token);
        this.router.navigate(['/events']);
      },
      error: err => {
        console.error("Login error:", err);
        if (err.status === 401 || err.status === 400) {
          this.errorMessage = "Неверный email или пароль";
        } else {
          this.errorMessage = "Ошибка сервера";
        }
      }
    });
  }
}
