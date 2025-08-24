<script setup lang="ts">
interface User {
  user_id: number
  username: string
  display_name: string
}

const user = useState<User>("user")
</script>

<template>
  <nav class="px-4 py-3 flex items-center space-x-8">
    <div class="text-lg font-bold">
      SheetHappens
    </div>
    <ul class="flex space-x-4">
      <li>
        <NuxtLink to="/characters">Characters</NuxtLink>
      </li>
      <li v-if="user">
        <NuxtLink to="/logout">Logout</NuxtLink>
      </li>
    </ul>
  </nav>
  <UApp>
    <NuxtPage />
  </UApp>
</template>
