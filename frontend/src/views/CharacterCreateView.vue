<template>
  <div>
    <h1>Create New Character</h1>
    <form @submit.prevent="createCharacter">
      <p><strong>Name:</strong> <input v-model="character.name" placeholder="name" /></p>
      <p><strong>Level:</strong> <input v-model="character.level" placeholder="level" type="number" /></p>
      <p><strong>Health:</strong> <input v-model="character.health" placeholder="health" type="number" /></p>
      <p><strong>Mental Health:</strong> <input v-model="character.mentalhealth" placeholder="mental health" type="number"/></p>
      <p><strong>Race:</strong> <input v-model="character.race" placeholder="race" /></p>
      <p><strong>Gender:</strong> <input v-model="character.gender" placeholder="gender" /></p>
      <p><strong>Height:</strong> <input v-model="character.height" placeholder="height" /></p>
      <p><strong>Weight:</strong> <input v-model="character.weight" placeholder="weight" /></p>
      <p><strong>Dodge:</strong> <input v-model="character.dodge" placeholder="dodge" type="number" /></p>
      <button type="submit">Create</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from "vue-router";

const character = ref<any>({});
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
          name: character.value.name,
          level: Number(character.value.level),
          health: Number(character.value.health),
          mentalhealth: Number(character.value.mentalhealth),
          race: character.value.race,
          gender: character.value.gender,
          height: character.value.height,
          weight: character.value.weight,
          dodge: Number(character.value.dodge),
        }),
      });
      if (response.ok) {
        router.push("/characters");
      } else {
        console.error("Character creation failed");
      }
  } catch (error) {
    console.error('Error creating character:', error);
  }
};
</script>
