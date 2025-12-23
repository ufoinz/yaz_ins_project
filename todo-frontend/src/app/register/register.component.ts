import { Component } from '@angular/core';
import { FormsModule } from "@angular/forms";
import { NgIf } from '@angular/common';
import { TodoService } from '../services/todo.service';
import { Router, RouterLink } from '@angular/router';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [FormsModule, NgIf, RouterLink],
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent {
  name = '';
  email = '';
  password = '';
  errorMessage: string = '';

  constructor(private todoServise: TodoService,  private router: Router) {}

  onRegister() {
    this.errorMessage = '';

    this.todoServise.register({
      name: this.name,
      email: this.email,
      password: this.password
    }).subscribe({
      next: res => {
        this.router.navigate(['/login']);
      },
      error: err => {
        console.error("Register error:", err);

        if (err.status === 409) {
          this.errorMessage = "Такой email уже зарегистрирован";
        } else if (err.status === 400) {
          this.errorMessage = "Некорректные данные";
        } else {
          this.errorMessage = "Ошибка сервера";
        }
      }
    });
  }
}
