
import { useUserStore } from '@/store/user'

export async function authorizedFetch(url, options = {}) {
  const userStore = useUserStore();

  const defaultHeaders = {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${userStore.accessToken}` 
  };
  const finalOptions = {
    ...options,
    headers: {
      ...defaultHeaders,
      ...(options.headers || {})
    },
    credentials: 'include',
  };

  const response = await fetch(url, finalOptions);

  if (response.status === 401) {
    const refreshRes = await fetch('http://localhost:8888/refreshToken', {
        method: 'POST',
        credentials: 'include', 
    });

    if (refreshRes.ok) {
        const newToken = await refreshRes.json();
        userStore.setAccessToken(newToken.accessToken); 
        // 重新发送原请求
        return await authorizedFetch(url, options);
    } else {
        // 无法续期，跳转登录页
        router.push('/login');
        return;
    }
 }
  
  return response;
}


