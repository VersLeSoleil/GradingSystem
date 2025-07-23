// stores/user.js
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const accessToken = ref(null)  // 只存在内存中
  const userInfo = ref(null)

  function setAccessToken(token) {
    accessToken.value = token
  }

  function setUserInfo(info) {
    userInfo.value = info
  }

  function logout() {
    accessToken.value = null
    userInfo.value = null
  }

  return {
    accessToken,
    userInfo,
    setAccessToken,
    setUserInfo,
    logout
  }
})
