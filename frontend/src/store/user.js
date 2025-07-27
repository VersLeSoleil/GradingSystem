import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const accessToken = ref(null)
  const userInfo = ref(null)

  function setAccessToken(token) {
    accessToken.value = token
  }

  function setUserInfo(info) {
    userInfo.value = info
  }

  async function logout() {
    try {
      await fetch('http://localhost:8888/logout', {
        method: 'POST',
        credentials: 'include',
      })
    } catch (e) {
      console.error('注销失败（可忽略）：', e)
    }

    accessToken.value = null
    userInfo.value = null
  }

  return {
    accessToken,
    userInfo,
    setAccessToken,
    setUserInfo,
    logout,
  }
}, {
  persist: true  
})
