<script lang="ts" setup>

const api = useRuntimeConfig().public.apiUrl
const headers = useRequestHeaders(["cookie"])
const route = useRoute()

interface Character {
  ID: number
  system_id: number
  name: string
  data: any
}

const { data: character, pending, error } = await useFetch<Character>(`${api}/characters/${route.params.id}`, {
  headers,
  method: "GET",
  credentials: "include",
})
</script>

<template>
  <div v-if="character">
    <h1>{{ character.name }}</h1>
    <div v-if="pending">Loading...</div>
    <div v-else-if="error">Error loading data</div>
    <div v-else>
      <DynamicRenderer :data="character.data" />
    </div>
  </div>
</template>
