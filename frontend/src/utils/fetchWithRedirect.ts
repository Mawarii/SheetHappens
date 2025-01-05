import { useRouter } from 'vue-router';

const fetchWithRedirect = async (url: string, options = {}) => {
  const router = useRouter();

  try {
    const response = await fetch(url, {
      ...options,
      credentials: 'include',
    });

    if (response.status === 401) {
      router.push('/');
      throw new Error('Unauthorized');
    }

    return response;
  } catch (error) {
    console.error('Fetch error:', error);
    throw error;
  }
};

export default fetchWithRedirect;
