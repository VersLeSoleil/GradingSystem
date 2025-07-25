<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const username = ref('');
const password = ref('');
const confirmPassword = ref('');
const showForm = ref(false)

onMounted(() => {
  setTimeout(() => {
    showForm.value = true
  }, 200)
})

const gotoLogin = () => {
  router.push('/login');
}
async function registerUser() {
  console.log('registerUser 被调用');
  if (!username.value || !password.value) {
    alert('请输入用户名和密码');
    return;
  }
  if (password.value !== confirmPassword.value) {
    alert('两次输入的密码不一致');
    return;
  }

  try {
    const endpoint = 'http://localhost:8888/register';
    const response = await fetch(endpoint, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        user_name: username.value,
        password: password.value,
      }),
      credentials: 'include',
    });

    const result = await response.json();
    if (response.ok) {
      alert('注册成功！');
      router.push('/login');
    } else {
      if (result.error && result.error.includes('Duplicate entry')) {
        alert('该用户名已被注册，请更换用户名');
      } else {
        alert(result.error || '注册失败！');
      }
    }
  } catch (error) {
    alert(error.message);
  }
}
</script>

<template>
  <div id="login-container">
    <br>
    <br>
    <transition name="slide-up">
      <div class="wrapper" v-if="showForm">
        <form @submit.prevent="loginBrungle">
          <h1>创建一个账号</h1>
          <div class="input-box">
            <i class="fas fa-envelope icon"></i>
            <input id="usernameInput" v-model="username" type="text" required />
            <label>用户名</label>
          </div>
          <div class="input-box">
            <input v-model="password" type="password" required />
            <label>密码</label>
          </div>
          <div class="input-box">
            <input v-model="confirmPassword" type="password" required />
            <label>确认密码</label>
          </div>
          <button class="btn" type="submit" @click="registerUser">注册</button>
          <div class="signup-link">
            <p>已有账号？ <a href="#" @click.prevent="gotoLogin">立即登录</a></p>
          </div>
        </form>
      </div>
    </transition>
  </div>
  
</template>



<style scoped>
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/all.min.css');
* {
  font-family: "Poppins", sans-serif;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
:global(body) {
  min-height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
  background: url('@/assets/background.png') center center/cover no-repeat;
  background-size: cover;
  background-position: center;
  margin: 0;
  padding: 0;
}
#app {
  width: 100vw;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}
#login-container {
  min-height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1000;
  background: none;
}
.title {
  font-size: 2em;
  color: #fff;
  text-align: center;
  margin-bottom: 20px;
}
.wrapper {
  width: 400px;
  height: 450px;
  background: rgba(112, 183, 230, 0.2);
  border: 2px solid rgba(255, 255, 255, 0.5);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  backdrop-filter: blur(15px);
  margin: 0 auto;
}
.wrapper:hover {
  box-shadow: 0 0 40px rgba(255,255,255,0.5);
  background: rgba(112, 183, 230, 0.2);
}
.wrapper h1,
.login-title {
  font-size: 2.2em;
  color: #2196f3;
  text-align: center;
  font-family: 'Segoe UI', 'Poppins', 'Arial', sans-serif;
  font-weight: 900;
  letter-spacing: 2.5px;
  margin-bottom: 18px;
  text-shadow: 0 4px 16px rgba(33,150,243,0.18), 0 1px 0 #fff;
  background: linear-gradient(90deg, #2196f3 30%, #4fc3f7 70%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
.wrapper .input-box {
  position: relative;
  width: 310px;
  margin: 30px 0;
  border-bottom: 2px solid #fff;
}
.wrapper .input-box input {
  width: 100%;
  height: 50px;
  background: transparent;
  outline: none;
  border: none;
  font-size: 1em;
  color: #fff;
  padding: 0 40px 0 5px;
}
.wrapper .input-box label {
  position: absolute;
  top: 50%;
  left: 5px;
  transform: translateY(-50%);
  font-size: 1em;
  color: #fff;
  pointer-events: none;
  transition: 0.5s;
}
.wrapper .input-box input:focus ~ label,
.wrapper .input-box input:valid ~ label {
  top: -5px;
}
.wrapper .input-box .icon {
  position: absolute;
  right: 8px;
  color: #fff;
  font-size: 1.2em;
  line-height: 57px;
}
.wrapper .row {
  margin: -15px 0 15px;
  font-size: 0.9em;
  color: #fff;
  display: flex;
  justify-content: space-between;
}
.wrapper .row label {
  display: flex;
  align-items: center;
  gap: 5px;
}
.wrapper .row a {
  color: #fff;
  text-decoration: none;
}
.wrapper .options a:hover {
  text-decoration: underline;
}
.wrapper .btn {
  width: 100%;
  height: 40px;
  background: #fff;
  outline: none;
  border: none;
  border-radius: 40px;
  cursor: pointer;
  font-size: 1em;
  font-weight: 500;
  color: #000;
  margin-top: 10px;
}
.btn:hover {
  background: #ffffea;
}
.wrapper .signup-link {
  font-size: 0.9em;
  color: #fff;
  text-align: center;
  margin: 25px 0 10px;
}
.wrapper .signup-link a {
  color: #fff;
  text-decoration: none;
  font-weight: 600;
}
.wrapper .signup-link a:hover {
  text-decoration: underline;
}
@media (max-width: 360px) {
  .wrapper {
    width: 100%;
    height: 100vh;
    border: none;
    border-radius: 0px;
  }
  .wrapper .input-box {
    width: 290px;
  }
  #img {
    z-index: -90;
  }
}
</style>

<style>
.slide-up-enter-active {
  transition: all 0.8s cubic-bezier(0.23, 1, 0.32, 1);
}
.slide-up-enter-from {
  opacity: 0;
  transform: translateY(60px);
}
.slide-up-enter-to {
  opacity: 1;
  transform: translateY(0);
}
</style>
