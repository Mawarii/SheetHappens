<template>
  <div>
    <h1>Create New Character</h1>
    <form @submit.prevent="createCharacter">
      <p><strong>Name:</strong> <input v-model="name" placeholder="name" /></p>
      <p><strong>Gender:</strong> <input v-model="gender" placeholder="gender" /></p>
      <p><strong>Level:</strong> <input v-model="level" placeholder="level" type="number" /></p>
      <button type="submit">Create</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from "vue-router";

const name = ref<string>("");
const gender = ref<string>("");
const level = ref<number>(0);
const router = useRouter();

const createCharacter = async () => {
try {
    const response = await fetch("http://localhost:3000/api/characters", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: name.value,
          gender: gender.value,
          level: Number(level.value),
        }),
      });
      if (response.ok) {
        router.push("/characters");
      } else {
        console.error("Login failed");
      }
  } catch (error) {
    console.error('Error fetching characters:', error);
  }
};
</script>
