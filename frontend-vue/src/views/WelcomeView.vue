<script setup lang="ts">
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'

const auth = useAuthStore()
const router = useRouter()

onMounted(() => {
  if (!auth.justRegistered) router.replace('/register')
})

function goToLogin() {
  auth.clearRegistered()
  window.location.href = 'http://localhost:4200/login'
}
</script>

<template>
  <div class="page">
    <div class="card">
      <div class="check-circle">
        <svg viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
          <circle cx="26" cy="26" r="25" stroke="#22c55e" stroke-width="2"/>
          <path d="M14 26l8 8 16-16" stroke="#22c55e" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </div>

      <h1>ยินดีต้อนรับ! 🎉</h1>
      <p class="username">{{ auth.registeredUsername }}</p>
      <p class="msg">สมัครสมาชิกเรียบร้อยแล้ว<br>กรุณาเข้าสู่ระบบเพื่อใช้งาน</p>

      <button id="welcome-login-btn" class="btn-primary" @click="goToLogin">
        เข้าสู่ระบบด้วย Angular →
      </button>

      <div class="badge">
        <span class="dot green"></span> บัญชีถูกบันทึกโดย C# Backend
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
  background: radial-gradient(ellipse at 40% 30%, #052e16 0%, #021712 50%, #000 100%);
  padding: 1.5rem;
}

.card {
  background: rgba(255,255,255,0.04);
  border: 1px solid rgba(34,197,94,0.15);
  border-radius: 24px;
  padding: 3rem 2.5rem;
  width: 100%;
  max-width: 400px;
  text-align: center;
  backdrop-filter: blur(24px);
  box-shadow: 0 25px 80px rgba(0,0,0,0.6), 0 0 60px rgba(34,197,94,0.05);
  animation: fadeUp 0.5s ease;
}

@keyframes fadeUp {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}

.check-circle {
  width: 80px;
  height: 80px;
  margin: 0 auto 1.5rem;
  animation: pop 0.5s ease;
}
@keyframes pop {
  0%   { transform: scale(0.5); opacity: 0; }
  80%  { transform: scale(1.1); }
  100% { transform: scale(1); opacity: 1; }
}

h1 { font-size: 1.8rem; font-weight: 700; color: white; margin: 0 0 0.5rem; }

.username {
  font-size: 1.1rem;
  font-weight: 600;
  color: #4ade80;
  background: rgba(34,197,94,0.1);
  display: inline-block;
  padding: 0.3rem 1rem;
  border-radius: 20px;
  border: 1px solid rgba(34,197,94,0.2);
  margin-bottom: 1rem;
}

.msg { color: rgba(255,255,255,0.55); font-size: 0.9rem; line-height: 1.6; margin-bottom: 2rem; }

.btn-primary {
  width: 100%;
  padding: 0.85rem;
  background: linear-gradient(135deg, #16a34a, #22c55e);
  color: white;
  font-size: 1rem;
  font-weight: 600;
  border: none;
  border-radius: 12px;
  cursor: pointer;
  transition: opacity 0.2s, transform 0.15s;
  margin-bottom: 1.5rem;
}

.btn-primary:hover { opacity: 0.9; transform: translateY(-1px); }

.badge {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  justify-content: center;
  font-size: 0.75rem;
  color: rgba(255,255,255,0.35);
}

.dot { width: 7px; height: 7px; border-radius: 50%; display: inline-block; }
.dot.green { background: #22c55e; }
</style>
