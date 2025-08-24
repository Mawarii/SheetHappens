<script setup lang="ts">
import type { FormError, FormSubmitEvent } from '@nuxt/ui'

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
const validate = (state: any): FormError[] => {
  const errors = []
  if (!state.username) errors.push({ name: 'username', message: 'Required' })
  if (!state.password) errors.push({ name: 'password', message: 'Required' })
  return errors
}
</script>

<template>
  <UForm
    :validate="validate"
    :state="state"
    class="space-y-4"
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
