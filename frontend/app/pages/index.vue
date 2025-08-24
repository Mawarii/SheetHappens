<script setup lang="ts">
const api = useRuntimeConfig().public.apiUrl

const state = reactive({
  username: undefined,
  password: undefined
})

async function onSubmit() {
  try {
    const res = await $fetch(`${api}/auth/login`, {
      method: "POST",
      credentials: "include",
      body: {
        username: state.username,
        password: state.password,
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
  <UForm
    :state="state"
    class="space-y-4 flex flex-col items-center"
    @submit="onSubmit"
  >
    <UFormField
      label="Username"
      name="username"
    >
      <UInput v-model="state.username" />
    </UFormField>

    <UFormField
      label="Password"
      name="password"
    >
      <UInput
        v-model="state.password"
        type="password"
      />
    </UFormField>

    <UButton type="submit">
      Submit
    </UButton>
  </UForm>
</template>
