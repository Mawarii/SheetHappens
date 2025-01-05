<template>
  <div class="dashboard">
    <h1>Characters <button class="add-btn" @click="createCharacter()" type="button"><OhVueIcon name="px-plus" scale="1.5" /></button></h1>
    <div v-if="characters" class="character-list">
      <div v-for="char in characters" :key="char['id']" @click="goToCharacterDetail(char['id'])" class="character-item">
        <span class="char-name">{{ char['name'] }}</span>
        <button class="delete-btn" @click.stop="confirmDelete(char['id'])" type="button">
          <OhVueIcon name="bi-trash-fill" />
        </button>
      </div>
    </div>
    <div v-else>
      <span>No Characters available!</span>
    </div>
  </div>

  <div v-if="showModal" class="modal-overlay">
    <div class="modal">
      <h2>Are you sure you want to delete this character?</h2>
      <div class="modal-actions">
        <button @click="deleteCharacter" class="confirm-btn">Yes, delete</button>
        <button @click="cancelDelete" class="cancel-btn">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import fetchWithRedirect from '@/utils/fetchWithRedirect';
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { OhVueIcon, addIcons } from "oh-vue-icons";
import { BiTrashFill, PxPlus } from "oh-vue-icons/icons";

addIcons(BiTrashFill, PxPlus)

const characters = ref<any | null>(null);
const showModal = ref(false);
const characterToDelete = ref<number | null>(null);
const router = useRouter();

const fetchCharacters = async () => {
  try {
    const res = await fetchWithRedirect("http://localhost:3000/api/characters", {
      method: "GET",
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

const confirmDelete = (id: number) => {
  characterToDelete.value = id;
  showModal.value = true;
};

const cancelDelete = () => {
  showModal.value = false;
  characterToDelete.value = null;
};

const deleteCharacter = async () => {
  try {
    if (characterToDelete.value !== null) {
      const res = await fetchWithRedirect(`http://localhost:3000/api/characters/${characterToDelete.value}`, {
        method: "DELETE",
      });
      if (res.ok) {
        showModal.value = false;
        characterToDelete.value = null;
        fetchCharacters();
      }
    }
  } catch (error) {
    console.error('Error deleting character:', error);
  }
};

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
  padding: 0;
  background-color: #448e47;
  color: rgb(255, 255, 255);
  border-radius: 50%;
  width: 40px;
  height: 40px;
}

.add-btn:hover {
  background-color: #397e3c;
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
}

.delete-btn:hover {
  background-color: #c0392b;
}

/* Modal Overlay */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

/* Modal */
.modal {
  background-color: rgb(0, 0, 0);
  padding: 20px;
  border-radius: 8px;
  text-align: center;
  width: 300px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.modal h2 {
  margin-bottom: 20px;
}

/* Actions for the modal */
.modal-actions {
  display: flex;
  justify-content: space-between;
}

.confirm-btn, .cancel-btn {
  padding: 8px 16px;
  cursor: pointer;
  border: none;
  border-radius: 4px;
  font-size: 14px;
}

.confirm-btn {
  background-color: #FF4136;
  color: white;
}

.cancel-btn {
  background-color: #636363;
  color: white;
}

.dashboard {
  width: 100%;
}
</style>
