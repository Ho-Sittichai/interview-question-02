<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const auth = useAuthStore()

const username = ref('')
const password = ref('')
const confirmPassword = ref('')
const errorMsg = ref('')
const successMsg = ref('')
const loading = ref(false)

async function handleRegister() {
  errorMsg.value = ''
  successMsg.value = ''

  if (!username.value || !password.value) {
    errorMsg.value = 'กรุณากรอกข้อมูลให้ครบถ้วน'
    return
  }
  if (password.value !== confirmPassword.value) {
    errorMsg.value = 'รหัสผ่านไม่ตรงกัน'
    return
  }
  if (password.value.length < 6) {
    errorMsg.value = 'รหัสผ่านต้องมีอย่างน้อย 6 ตัวอักษร'
    return
  }

  loading.value = true
  try {
    // C# Backend (Port 5001) — Register Only
    await axios.post('http://localhost:5001/api/auth/register', {
      username: username.value,
      password: password.value
    })
    auth.setRegistered(username.value)
    successMsg.value = 'สมัครสมาชิกสำเร็จ! กำลังพาคุณไปหน้ายินดีต้อนรับ...'
    setTimeout(() => router.push('/welcome'), 1200)
  } catch (err: any) {
    errorMsg.value = err.response?.data?.error || 'เกิดข้อผิดพลาด กรุณาลองใหม่'
  } finally {
    loading.value = false
  }
}

function goToLogin() {
  // Redirect to Angular Login App (Port 4200)
  window.location.href = 'http://localhost:4200/login'
}
</script>

<template>
  <div class="page">
    <div class="card">
      <div class="logo">
        <div class="logo-icon">✦</div>
        <h1 class="brand">AuthSplit</h1>
      </div>

      <div class="header">
        <h2>สร้างบัญชีใหม่</h2>
        <p class="subtitle">Register via C# Backend · Port 5001</p>
      </div>

      <form @submit.prevent="handleRegister" id="register-form">
        <div class="field">
          <label for="reg-username">Username</label>
          <input
            id="reg-username"
            v-model="username"
            type="text"
            placeholder="กรอก username ของคุณ"
            autocomplete="username"
          />
        </div>

        <div class="field">
          <label for="reg-password">Password</label>
          <input
            id="reg-password"
            v-model="password"
            type="password"
            placeholder="อย่างน้อย 6 ตัวอักษร"
            autocomplete="new-password"
          />
        </div>

        <div class="field">
          <label for="reg-confirm">ยืนยัน Password</label>
          <input
            id="reg-confirm"
            v-model="confirmPassword"
            type="password"
            placeholder="กรอก password อีกครั้ง"
            autocomplete="new-password"
          />
        </div>

        <div v-if="errorMsg" class="alert error" role="alert">{{ errorMsg }}</div>
        <div v-if="successMsg" class="alert success" role="status">{{ successMsg }}</div>

        <button id="register-btn" type="submit" class="btn-primary" :disabled="loading">
          <span v-if="loading" class="spinner"></span>
          <span v-else>สมัครสมาชิก</span>
        </button>
      </form>

      <div class="divider"><span>หรือ</span></div>

      <button id="goto-login-btn" class="btn-secondary" @click="goToLogin">
        มีบัญชีอยู่แล้ว? เข้าสู่ระบบ →
      </button>

      <div class="badge">
        <span class="dot green"></span> C# Register Backend
        <span class="separator">·</span>
        <span class="dot blue"></span> Vue Register App
      </div>
    </div>
  </div>
</template>

<style scoped>
.page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: radial-gradient(ellipse at 60% 20%, #1a1060 0%, #0a0718 60%, #000 100%);
  padding: 1.5rem;
}

.card {
  background: rgba(255,255,255,0.04);
  border: 1px solid rgba(255,255,255,0.1);
  border-radius: 24px;
  padding: 2.5rem;
  width: 100%;
  max-width: 420px;
  backdrop-filter: blur(24px);
  box-shadow: 0 25px 80px rgba(0,0,0,0.6), 0 0 0 1px rgba(255,255,255,0.05);
  animation: fadeUp 0.5s ease;
}

@keyframes fadeUp {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}

.logo {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  margin-bottom: 1.8rem;
}

.logo-icon {
  width: 38px;
  height: 38px;
  background: linear-gradient(135deg, #a78bfa, #7c3aed);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1rem;
  color: white;
}

.brand {
  font-size: 1.4rem;
  font-weight: 700;
  color: white;
  margin: 0;
  letter-spacing: -0.5px;
}

.header h2 {
  font-size: 1.6rem;
  font-weight: 700;
  color: white;
  margin: 0 0 0.3rem;
}

.subtitle {
  font-size: 0.78rem;
  color: rgba(167,139,250,0.8);
  margin: 0 0 1.8rem;
  font-family: monospace;
  letter-spacing: 0.03em;
}

.field {
  margin-bottom: 1.1rem;
}

label {
  display: block;
  font-size: 0.85rem;
  color: rgba(255,255,255,0.6);
  margin-bottom: 0.4rem;
  font-weight: 500;
}

input {
  width: 100%;
  padding: 0.75rem 1rem;
  background: rgba(255,255,255,0.06);
  border: 1px solid rgba(255,255,255,0.12);
  border-radius: 10px;
  color: white;
  font-size: 0.95rem;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
  box-sizing: border-box;
}

input::placeholder { color: rgba(255,255,255,0.25); }
input:focus {
  border-color: #7c3aed;
  box-shadow: 0 0 0 3px rgba(124,58,237,0.2);
}

.alert {
  padding: 0.7rem 1rem;
  border-radius: 8px;
  font-size: 0.88rem;
  margin-bottom: 1rem;
}

.alert.error { background: rgba(239,68,68,0.15); color: #fca5a5; border: 1px solid rgba(239,68,68,0.3); }
.alert.success { background: rgba(34,197,94,0.15); color: #86efac; border: 1px solid rgba(34,197,94,0.3); }

.btn-primary {
  width: 100%;
  padding: 0.85rem;
  background: linear-gradient(135deg, #7c3aed, #a855f7);
  color: white;
  font-size: 1rem;
  font-weight: 600;
  border: none;
  border-radius: 12px;
  cursor: pointer;
  transition: opacity 0.2s, transform 0.15s;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.btn-primary:hover:not(:disabled) { opacity: 0.9; transform: translateY(-1px); }
.btn-primary:disabled { opacity: 0.5; cursor: not-allowed; }

.spinner {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(255,255,255,0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }

.divider {
  text-align: center;
  position: relative;
  margin: 0.5rem 0;
  color: rgba(255,255,255,0.25);
  font-size: 0.8rem;
}
.divider::before, .divider::after {
  content: '';
  position: absolute;
  top: 50%;
  width: calc(50% - 1.5rem);
  height: 1px;
  background: rgba(255,255,255,0.1);
}
.divider::before { left: 0; }
.divider::after { right: 0; }

.btn-secondary {
  width: 100%;
  padding: 0.75rem;
  background: transparent;
  color: rgba(167,139,250,0.9);
  font-size: 0.9rem;
  font-weight: 500;
  border: 1px solid rgba(124,58,237,0.3);
  border-radius: 12px;
  cursor: pointer;
  transition: background 0.2s, border-color 0.2s;
  margin-top: 0.5rem;
}

.btn-secondary:hover { background: rgba(124,58,237,0.1); border-color: rgba(124,58,237,0.6); }

.badge {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  justify-content: center;
  margin-top: 1.5rem;
  font-size: 0.75rem;
  color: rgba(255,255,255,0.35);
}

.dot {
  width: 7px;
  height: 7px;
  border-radius: 50%;
  display: inline-block;
}

.dot.green { background: #22c55e; }
.dot.blue { background: #60a5fa; }
.separator { margin: 0 0.2rem; }
</style>
