<template>
  <div>
    <h1>Character Detail</h1>
    <form v-if="character">
    <p><strong>ID:</strong> {{ character.id }}</p>
    <p><strong>Name:</strong> {{ character.name }}</p>
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
      character: null as Character | null,
    };
  },
  async mounted() {
    const tokenStore = useTokenStore();
    const id = this.$route.params.id;
    console.log(id)
    try {
      const res = await fetch(`http://localhost:3000/api/characters/${id}`, {
        method: "GET",
        headers: {
          "Authorization": "Bearer " + tokenStore.token,
        },
      });
      const data = await res.json();
      this.character = data.character;
    } catch (error) {
      console.error('Error fetching characters:', error);
    }
  },
});
</script>
