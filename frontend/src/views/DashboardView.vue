<template>
  <div>
    <h1>Characters</h1>
    <form v-for="char in characters" :key="char.id">
      <button @click="goToCharacterDetail(char.id)" type="button">{{ char.name }}</button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useTokenStore } from '@/store';

interface Character {
  id: number;
  name: string;
}

export default defineComponent({
  data() {
    return {
      characters: null as Character[] | null,
    };
  },
  async mounted() {
    const tokenStore = useTokenStore();
    try {
      const res = await fetch("http://localhost:3000/api/characters", {
        method: "GET",
        headers: {
          "Authorization": "Bearer " + tokenStore.token,
        },
      });
      const data = await res.json();
      this.characters = data.characters;
    } catch (error) {
      console.error('Error fetching characters:', error);
    }
  },
  methods: {
    async goToCharacterDetail(id: number) {
      this.$router.push(`/edit/${id}`);
    },
  },
});
</script>
