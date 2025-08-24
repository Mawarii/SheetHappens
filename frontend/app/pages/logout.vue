<script setup lang="ts">
const api = useRuntimeConfig().public.apiUrl
const headers = useRequestHeaders(["cookie"])

const res = await $fetch(`${api}/auth/logout`, {
  headers,
  method: "GET",
  credentials: "include",
})
if (res) {
  navigateTo("/")
}
</script>

<template>
  <h1>Logout</h1>
</template>
