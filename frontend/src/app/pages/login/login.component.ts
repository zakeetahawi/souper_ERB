import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { AuthService } from '../../core/auth/auth.service';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    MatIconModule
  ],
  template: `
    <div class="login-container">
      <div class="login-card">
        <div class="login-header">
          <h1>ELKHAWAGA ERP</h1>
          <p>نظام إدارة الموارد المؤسسية</p>
        </div>
        
        <form [formGroup]="loginForm" (ngSubmit)="onSubmit()" class="login-form">
          <mat-form-field appearance="outline" class="full-width">
            <mat-label>اسم المستخدم</mat-label>
            <input matInput formControlName="username" placeholder="أدخل اسم المستخدم">
            <mat-error *ngIf="loginForm.get('username')?.hasError('required')">
              اسم المستخدم مطلوب
            </mat-error>
          </mat-form-field>

          <mat-form-field appearance="outline" class="full-width">
            <mat-label>كلمة المرور</mat-label>
            <input matInput type="password" formControlName="password" placeholder="أدخل كلمة المرور">
            <button mat-icon-button matSuffix (click)="togglePasswordVisibility()" type="button">
              <mat-icon>{{ hidePassword ? 'visibility' : 'visibility_off' }}</mat-icon>
            </button>
            <mat-error *ngIf="loginForm.get('password')?.hasError('required')">
              كلمة المرور مطلوبة
            </mat-error>
          </mat-form-field>

          <button 
            mat-raised-button 
            color="primary" 
            type="submit" 
            class="login-button"
            [disabled]="loginForm.invalid || isLoading">
            <mat-icon *ngIf="isLoading">hourglass_empty</mat-icon>
            {{ isLoading ? 'جاري تسجيل الدخول...' : 'تسجيل الدخول' }}
          </button>
        </form>

        <div class="login-footer">
          <p>نسيت كلمة المرور؟ <a href="#" (click)="forgotPassword($event)">اضغط هنا</a></p>
        </div>
      </div>
    </div>
  `,
  styles: [`
    .login-container {
      height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      padding: 20px;
    }

    .login-card {
      background: white;
      border-radius: 16px;
      box-shadow: 0 20px 40px rgba(0,0,0,0.1);
      padding: 40px;
      width: 100%;
      max-width: 400px;
      text-align: center;
    }

    .login-header {
      margin-bottom: 30px;
    }

    .login-header h1 {
      color: #333;
      font-size: 2rem;
      font-weight: 700;
      margin-bottom: 8px;
    }

    .login-header p {
      color: #666;
      font-size: 1rem;
      margin: 0;
    }

    .login-form {
      display: flex;
      flex-direction: column;
      gap: 20px;
    }

    .full-width {
      width: 100%;
    }

    .login-button {
      height: 48px;
      font-size: 1.1rem;
      font-weight: 600;
      border-radius: 8px;
      margin-top: 10px;
    }

    .login-footer {
      margin-top: 30px;
      padding-top: 20px;
      border-top: 1px solid #eee;
    }

    .login-footer p {
      color: #666;
      font-size: 0.9rem;
      margin: 0;
    }

    .login-footer a {
      color: #1976d2;
      text-decoration: none;
      font-weight: 500;
    }

    .login-footer a:hover {
      text-decoration: underline;
    }

    @media (max-width: 480px) {
      .login-card {
        padding: 30px 20px;
      }
      
      .login-header h1 {
        font-size: 1.5rem;
      }
    }
  `]
})
export class LoginComponent {
  loginForm: FormGroup;
  isLoading = false;
  hidePassword = true;

  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
    private router: Router
  ) {
    this.loginForm = this.fb.group({
      username: ['', [Validators.required]],
      password: ['', [Validators.required]]
    });
  }

  onSubmit(): void {
    if (this.loginForm.valid) {
      this.isLoading = true;
      
      this.authService.login(this.loginForm.value).subscribe({
        next: (response) => {
          if (response.success) {
            Swal.fire({
              icon: 'success',
              title: 'تم تسجيل الدخول بنجاح',
              text: 'مرحباً بك في نظام ELKHAWAGA ERP',
              timer: 2000,
              showConfirmButton: false
            }).then(() => {
              this.router.navigate(['/dashboard']);
            });
          }
        },
        error: (error) => {
          this.isLoading = false;
          let errorMessage = 'حدث خطأ أثناء تسجيل الدخول';
          
          if (error.error?.error) {
            errorMessage = error.error.error;
          }
          
          Swal.fire({
            icon: 'error',
            title: 'خطأ في تسجيل الدخول',
            text: errorMessage
          });
        },
        complete: () => {
          this.isLoading = false;
        }
      });
    }
  }

  togglePasswordVisibility(): void {
    this.hidePassword = !this.hidePassword;
  }

  forgotPassword(event: Event): void {
    event.preventDefault();
    Swal.fire({
      icon: 'info',
      title: 'نسيت كلمة المرور',
      text: 'سيتم إرسال رابط إعادة تعيين كلمة المرور إلى بريدك الإلكتروني',
      showCancelButton: true,
      confirmButtonText: 'إرسال',
      cancelButtonText: 'إلغاء'
    });
  }
} 