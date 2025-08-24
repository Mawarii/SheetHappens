<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'

const items = ref<NavigationMenuItem[][]>([
  [
    {
      label: 'SheetHappens',
      icon: 'i-lucide-book-open',
    },
    {
      label: 'Characters',
      icon: 'i-lucide-database',
      to: '/characters'
    }
  ],
  [
    {
      label: 'Logout',
      icon: 'i-lucide-log-out',
      to: '/logout'
    }
  ]
])

interface User {
  user_id: number
  username: string
  display_name: string
}

const user = useState<User>("user")

</script>

<template>
  <UNavigationMenu
    color="neutral"
    variant="link"
    highlight
    :items="items"
    class="w-full"
  />
  <slot />
</template>
