<script lang="ts" setup>
const api = useRuntimeConfig().public.apiUrl
const headers = useRequestHeaders(["cookie"])

const { data: chars } = useFetch(api + "/characters", {
  headers,
  method: "GET",
  credentials: "include",
})
</script>

<template>
  <h1>Characters</h1>
  <article
    v-for="char in chars"
    v-bind:key="char.id"
  >
    {{ char.name }}
  </article>
</template>
