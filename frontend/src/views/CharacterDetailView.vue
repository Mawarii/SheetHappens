<template>
  <div>
    <h1>Character Details</h1>
    <form v-if="character">
      <p><strong>ID:</strong> {{ character['id'] }}</p>
      <p><strong>Name:</strong> {{ character['name'] }}</p>
      <p><strong>Gender:</strong> {{ character['gender'] }}</p>
      <p><strong>Level:</strong> {{ character['level'] }}</p>
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
