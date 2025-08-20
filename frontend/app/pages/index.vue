<script lang="ts" setup>
const api = useRuntimeConfig().public.apiUrl

const headers = useRequestHeaders(["cookie"])

const { data: user } = await useFetch(api + "/auth/info", {
  headers,
  credentials: "include",
})


const username = ref("");
const password = ref("");

async function OnSubmit() {
  try {
    const res = await $fetch(api + "/auth/login", {
      method: "POST",
      credentials: "include",
      body: {
        username: username.value,
        password: password.value,
      },
    })
    if (res) {
      navigateTo("/characters")
    }
  } catch (e) {
    console.error("Error during login:", e)
  }
}
</script>

<template>
  <main>
    <div class="container">
      <h1>Login {{ user.username }}</h1>
      <form @submit.prevent="OnSubmit">
        <input
          v-model="username"
          type="text"
          placeholder="Username"
          aria-label="Login"
          autoComplete="username"
          required
        />
        <input
          v-model="password"
          type="password"
          placeholder="Password"
          aria-label="Password"
          autoComplete="current-password"
          required
        />
        <button type="submit">Login</button>
      </form>
    </div>
  </main>
</template>

<style scoped>
main {
  max-width: 600px;
  margin: 4rem auto;
}
</style>
