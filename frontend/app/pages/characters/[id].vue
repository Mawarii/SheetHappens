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

const { data: character, pending } = await useFetch<Character>(`${api}/characters/${route.params.id}`, {
  headers,
  method: "GET",
  credentials: "include",
})

const edit = ref(false)

const saveCharacter = async () => {
  await $fetch(`${api}/characters/${route.params.id}`, {
    headers,
    method: "PUT",
    credentials: "include",
    body: JSON.stringify(character.value),
  })
}
</script>

<template>
  <div v-if="character">
    <h1>{{ character.name }}</h1>
    <USwitch
      v-model="edit"
      label="Edit"
    />
    <div v-if="pending">Loading...</div>
    <div v-else>
      <DynamicRenderer
        :data="character.data"
        :editable="edit"
      />
      <UButton
        v-if="edit"
        @click="saveCharacter"
      >
        Save
      </UButton>
    </div>
  </div>
</template>
