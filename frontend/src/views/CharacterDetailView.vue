<template>
  <div>
    <h1>Character Details</h1>
    <form v-if="character">
      <p><strong>Name:</strong> {{ character['name'] }}</p>
      <p v-if="character['level']"><strong>Level:</strong> {{ character['level'] }}</p>
      <p v-if="character['health']"><strong>Health:</strong> {{ character['health'] }}</p>
      <p v-if="character['mentalhealth']"><strong>Mental Health:</strong> {{ character['mentalhealth'] }}</p>
      <p v-if="character['race']"><strong>Race:</strong> {{ character['race'] }}</p>
      <p v-if="character['gender']"><strong>Gender:</strong> {{ character['gender'] }}</p>
      <p v-if="character['height']"><strong>Height:</strong> {{ character['height'] }}</p>
      <p v-if="character['weight']"><strong>Weight:</strong> {{ character['weight'] }}</p>
      <p v-if="character['dodge']"><strong>Dodge:</strong> {{ character['dodge'] }}</p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';

const character = ref<any | null>(null);
const route = useRoute();

const fetchCharacter = async () => {
try {
    const id = route.params.id;
    const res = await fetch(`http://localhost:3000/api/characters/${id}`, {
        method: "GET",
        credentials: 'include',
      });
      const data = await res.json();
      character.value = data.character;
  } catch (error) {
    console.error('Error fetching characters:', error);
  }
};

onMounted(fetchCharacter);
</script>
