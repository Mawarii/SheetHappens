<template>
  <div>
    <h1>Create New Character</h1>
    <form @submit.prevent="createCharacter">
      <p><strong>Name:</strong> <input v-model="name" placeholder="name" /></p>
      <p><strong>Level:</strong> <input v-model="level" placeholder="level" type="number" /></p>
      <p><strong>Health:</strong> <input v-model="health" placeholder="health" type="number" /></p>
      <p><strong>Mental Health:</strong> <input v-model="mentalhealth" placeholder="mental health" type="number"/></p>
      <p><strong>Race:</strong> <input v-model="race" placeholder="race" /></p>
      <p><strong>Gender:</strong> <input v-model="gender" placeholder="gender" /></p>
      <p><strong>Height:</strong> <input v-model="height" placeholder="height" /></p>
      <p><strong>Weight:</strong> <input v-model="weight" placeholder="weight" /></p>
      <p><strong>Dodge:</strong> <input v-model="dodge" placeholder="dodge" type="number" /></p>
      <button type="submit">Create</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from "vue-router";

const name = ref<string>("");
const level = ref<number>(0);
const health = ref<number>(0);
const mentalhealth = ref<number>(0);
const race = ref<string>("");
const gender = ref<string>("");
const height = ref<string>("");
const weight = ref<string>("");
const dodge = ref<number>(0);

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
          level: Number(level.value),
          health: Number(health.value),
          mentalhealth: Number(mentalhealth.value),
          race: race.value,
          gender: gender.value,
          height: height.value,
          weight: weight.value,
          dodge: Number(dodge.value),
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
