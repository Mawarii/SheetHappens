<script lang="ts" setup>
const api = useRuntimeConfig().public.apiUrl
const headers = useRequestHeaders(["cookie"])

interface Character {
  ID: number
  system_id: number
  name: string
  data: JSON
}

const { data: chars } = await useFetch<Character[]>(`${api}/characters`, {
  headers,
  method: "GET",
  credentials: "include",
})
</script>

<template>
  <h1>Characters</h1>
  <div
    v-if="chars"
    class="grid grid-cols-2 gap-3"
  >
    <UButton
      v-for="char in chars"
      @click="navigateTo(`/characters/${char.ID}`)"
    >
      {{ char.name }}
    </UButton>
  </div>
  <div v-else>
    <h1>No Characters</h1>
  </div>
</template>
