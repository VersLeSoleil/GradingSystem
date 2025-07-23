<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import {useUserStore} from '@/store/user';
//import {axios} from 'axios';

const username = ref('');
const password = ref('');
const router = useRouter();
const userStore = useUserStore();

function forgotman() {
  alert('Please email ssnigdhasiraz22@sirhenryfloyd.co.uk to request a password reset');
}

async function loginBrungle() {
  try {
    let endpoint = 'http://localhost:8888/login';
    let method = 'POST';
    let requestBody = {
      user_name: username.value,
      password: password.value
    };
      // 发送请求到后端
    const response = await fetch(endpoint, {
      method: method,
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(requestBody),
      credentials: 'include', 
    });

    // 调试：打印响应状态码和响应内容
    console.log('Response Status:', response.status);
    console.log('Response Headers:', response.headers);

    // 解析响应
    const result = await response.json();
    console.log('Response Data:', result);
    userStore.print();
    // 处理成功与否
    if (response.ok) {
      userStore.setAccessToken(result.access_token)
      userStore.setUserInfo(result.user)
      userStore.print();
      alert('登录成功！');
      router.push('/home'); 
    } else {
      // 错误处理
      alert(result.error || '登录失败！');
    }
  } catch (error) {
    alert(error.message) 
  }
}

//注册
const registerDialogVisible = ref(false)
const registerUsername = ref('')
const registerPassword = ref('')
const confirmPassword = ref('')
const sex = ref('')
const birthday=ref('')



function openRegisterDialog() {
  registerDialogVisible.value = true
}

function closeRegisterDialog() {
  registerDialogVisible.value = false
  registerUsername.value = ''
  registerPassword.value = ''
  confirmPassword.value = ''
}

async function registerUser() {
  if (!registerUsername.value || !registerPassword.value) {
    alert('请输入用户名和密码')
    return
  }
  if (registerPassword.value !== confirmPassword.value) {
    alert('两次输入的密码不一致')
    return
  }
  try {
    let endpoint = 'http://localhost:8888/register';
    let method = 'POST';
    let requestBody = {
      register_username: registerUsername.value,
      register_password: registerPassword.value
    };
      // 发送请求到后端
    const response = await fetch(endpoint, {
      method: method,
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(requestBody),
      credentials: 'include', 
    });

    // 调试：打印响应状态码和响应内容
    console.log('Response Status:', response.status);
    console.log('Response Headers:', response.headers);

    // 解析响应
    const result = await response.json();
    console.log('Response Data:', result);
    userStore.print();
    // 处理成功与否
    if (response.ok) {
      userStore.setAccessToken(result.access_token)
      userStore.setUserInfo(result.user)
      userStore.print();
      alert('登录成功！');
      router.push('/home'); 
    } else {
      // 错误处理
      alert(result.error || '登录失败！');
    }
  } catch (error) {
    alert(error.message) 
  }
  alert(`注册成功！用户名: ${registerUsername.value}`)
  closeRegisterDialog()
}
</script>



<template>
  <div id="login-container">
    <h1 class="login-title">欢迎使用超声影像智能分级系统</h1>
    <br>
    <br>
      <div class="wrapper">
        <form @submit.prevent="loginBrungle">
          <h1>Login</h1>
          <div class="input-box">
            <i class="fas fa-envelope icon"></i>
            <input id="usernameInput" v-model="username" type="text" required />
            <label>Username</label>
          </div>
          <div class="input-box">
            <i class="fas fa-lock icon"></i>
            <input id="passwordInput" v-model="password" type="password" required />
            <label>Password</label>
          </div>
          <div class="row">
            <a href="#" @click.prevent="forgotman">Forgot password?</a>
          </div>
          <button class="btn" type="submit">Login</button>
          <div class="signup-link">
            <p>Don't have an account? <a href="#" @click.prevent="openRegisterDialog">Create one.</a></p>
          </div>
        </form>
      </div>
  </div>
  <el-dialog v-model="registerDialogVisible" title="注册账号" width="400px" :before-close="closeRegisterDialog">
    <div style="display: flex; flex-direction: column; gap: 16px; padding: 8px 0;">
      <el-input v-model="registerUsername" placeholder="请输入用户名" />
      <el-input v-model="registerPassword" placeholder="请输入密码" show-password />
      <el-input v-model="confirmPassword" placeholder="请再次输入密码" show-password />
      <el-input v-model="sex" placeholder="请" show-password />

      <div style="text-align: right;">
        <el-button @click="closeRegisterDialog">取消</el-button>
        <el-button type="primary" @click="registerUser">注册</el-button>
      </div>
    </div>
  </el-dialog>
  
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
