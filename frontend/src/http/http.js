import { useUserStore } from '@/stores/user'
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



  if (res.status === 401) {
    const refreshRes = await fetch('http://localhost:8888/refresh', {
        method: 'POST',
        credentials: 'include', // 关键，带上 refresh_token 的 Cookie
    });

    if (refreshRes.ok) {
        const newToken = await refreshRes.json();
        accessToken.value = newToken.access_token;

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


