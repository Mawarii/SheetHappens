<template>
  <div>
    <h1>Characters</h1>
    <form v-if="characters">
      <form v-for="char in characters" :key="char['id']">
        <button @click="goToCharacterDetail(char['id'])" type="button">{{ char['name'] }}</button>
      </form>
    </form>
    <form v-else>
      <span>No Characters available!</span>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const characters = ref<any | null>(null);
const router = useRouter();

const fetchCharacters = async () => {
  try {
    const res = await fetch("http://localhost:3000/api/characters", {
      method: "GET",
      credentials: 'include',
    });
    const data = await res.json();
    characters.value = data.characters;
  } catch (error) {
    console.error('Error fetching characters:', error);
  }
};

const goToCharacterDetail = (id: number) => {
  router.push(`/characters/${id}`);
};

onMounted(fetchCharacters);
</script>
