export default defineNuxtRouteMiddleware(async (to) => {
  const api = useRuntimeConfig().public.apiUrl
  const headers = useRequestHeaders(["cookie"])

  interface User {
    user_id: number
    username: string
    display_name: string
  }

  const user = useState<User>("user")

  const { data, error } = await useFetch<User>(`${api}/auth/info`, {
    headers,
    credentials: "include",
  })

  if (data.value) {
    user.value = data.value
  }

  if (error.value?.statusCode === 401 && to.path !== "/") {
    return navigateTo("/")
  }
  else if (data.value && to.path === "/") {
    return navigateTo("/characters")
  }

  return
})
