<template>
  <div>
    <h1>Characters 
      <button class="add-btn" @click="createCharacter()" type="button">+</button>
    </h1>
    <div v-if="characters" class="character-list">
      <div 
        v-for="char in characters" 
        :key="char['id']" 
        class="character-item"
        @click="goToCharacterDetail(char['id'])"
      >
        <span class="char-name">{{ char['name'] }}</span>
        <button class="delete-btn" @click.stop="deleteCharacter(char['id'])" type="button">
          <OhVueIcon name="bi-trash-fill" />
        </button>
      </div>
    </div>
    <div v-else>
      <span>No Characters available!</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { OhVueIcon, addIcons } from "oh-vue-icons";
import { BiTrashFill } from "oh-vue-icons/icons";

addIcons(BiTrashFill)

const characters = ref<any | null>(null);
const router = useRouter();

const fetchCharacters = async () => {
  try {
    const res = await fetch("http://localhost:3000/api/characters", {
      method: "GET",
      credentials: "include",
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

const deleteCharacter = async (id: number) => {
  try {
    const res = await fetch(`http://localhost:3000/api/characters/${id}`, {
      method: "DELETE",
      credentials: "include",
    });
    if (res.ok) {
      await fetchCharacters();
    }
  } catch (error) {
    console.error('Error deleting character:', error);
  }  
}

const createCharacter = () => {
  router.push("/characters/create");
}

onMounted(fetchCharacters);
</script>

<style scoped>
h1 {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.add-btn {
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  font-size: 24px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.add-btn:hover {
  background-color: #45a049;
}

.character-list {
  margin-top: 20px;
}

.character-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  margin-bottom: 12px;
  border: 1px solid #626262;
  border-radius: 8px;
  background-color: #000000;
  cursor: pointer;
  transition: box-shadow 0.3s, background-color 0.3s;
}

.character-item:hover {
  background-color: #252525;
}

.character-item:has(.delete-btn:hover) {
  background-color: #000000;
}

.char-name {
  font-size: 18px;
  color: #afafaf;
}

.delete-btn {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.delete-btn:hover {
  background-color: #c0392b;
}
</style>
