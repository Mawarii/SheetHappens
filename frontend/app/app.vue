<script setup lang="ts">
interface User {
  user_id: number
  username: string
  display_name: string
}

const user = useState<User | undefined>("user")

const api = useRuntimeConfig().public.apiUrl
const headers = useRequestHeaders(["cookie"])

async function logout() {
  try {
    const res = await $fetch(`${api}/auth/logout`, {
      headers,
      method: "GET",
      credentials: "include",
    })
    if (res) {
      await navigateTo("/")
      user.value = undefined
    }
  } catch (e) {
    console.error("Error during logout:", e)
  }
}
</script>

<template>
  <nav class="flex w-full max-w-full items-center justify-between px-4 py-2 mx-auto">
    <ul class="flex items-center gap-4">
      <li class="text-lg font-bold">SheetHappens</li>
      <li v-if="user">
        <UButton
          to="/characters"
          color="neutral"
          variant="outline"
          icon="i-lucide-files"
        >Characters</UButton>
      </li>
    </ul>
    <ul class="flex items-center">
      <li v-if="user">
        <UButton
          @click="logout"
          color="neutral"
          variant="outline"
          icon="i-lucide-log-out"
        />
      </li>
    </ul>
  </nav>
  <UApp>
    <NuxtPage />
  </UApp>
</template>
