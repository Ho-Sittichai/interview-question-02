import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, FormsModule, RouterModule],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  username = '';
  password = '';
  errorMsg = '';
  loading = false;

  constructor(private auth: AuthService, private router: Router) {}

  handleLogin(): void {
    this.errorMsg = '';

    if (!this.username || !this.password) {
      this.errorMsg = 'กรุณากรอก Username และ Password';
      return;
    }

    this.loading = true;
    this.auth.login(this.username, this.password).subscribe({
      next: () => {
        this.loading = false;
        this.router.navigate(['/welcome']);
      },
      error: (err) => {
        this.loading = false;
        this.errorMsg = err.error?.error ?? 'เกิดข้อผิดพลาด กรุณาลองใหม่';
      }
    });
  }

  goToRegister(): void {
    // Redirect to Vue Register App (Port 5173)
    window.location.href = 'http://localhost:5173/register';
  }
}
