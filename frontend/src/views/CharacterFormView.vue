<template>
  <div>
    <h1>{{ editMode ? 'Edit Character' : 'Create New Character' }}</h1>
    <form @submit.prevent="editMode ? updateCharacter() : createCharacter()">
      <p>Name: <input v-model="character.name" placeholder="name" /></p>
      <p>System: <select v-model="selectedSystem" @change="loadTemplate">
                  <option value="selfmade">Selfmade</option>
                  <option value="dnd5e">D&D 5e</option>
                  <option value="cthulhu">Call of Cthulhu</option>
                </select></p>

      <div v-for="(value, key) in character.data" :key="key">
        <label>{{ key }}:</label>
        <input v-model="character.data[key]" />
      </div>

      <div v-if="selectedSystem === 'selfmade'">
        <input v-model="newFieldKey" placeholder="Field name" />
        <input v-model="newFieldValue" placeholder="Field value" />
        <button type="button" @click="addField">Add Field</button>
      </div>

      <button type="submit">{{ editMode ? 'Save Changes' : 'Create Character' }}</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import fetchWithRedirect from '@/utils/fetchWithRedirect';
import { CoAral } from 'oh-vue-icons/icons';
import { ref, onMounted } from 'vue';
import { useRouter, useRoute } from "vue-router";

const character = ref<any>({});
const router = useRouter();
const route = useRoute();
const editMode = route.params.id;
const selectedSystem = ref('');
const templates = {
  dnd5e: {
    health: 10,
    mana: 5,
    strength: 15,
    dexterity: 14,
  },
  cthulhu: {
    sanity: 70,
    strength: 40,
    intelligence: 60,
  },
};

const loadTemplate = () => {
  if (selectedSystem.value === 'selfmade') {
    character.data = {};
  } else {
    character.value.data = { ...templates[selectedSystem.value] };
  }
  character.value.system = selectedSystem.value;
};

const newFieldKey = ref('');
const newFieldValue = ref('');

const addField = () => {
  if (newFieldKey.value) {
    if (!character.value.data) character.value.data = {};
    character.value.data[newFieldKey.value] = newFieldValue.value;
    newFieldKey.value = '';
    newFieldValue.value = '';
  }
};

const createCharacter = async () => {
  try {
    const response = await fetchWithRedirect("http://localhost:3000/api/characters", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(character.value),
    });
    if (response.ok) {
      const data = await response.json()
      console.log(data);
    } else {
      console.error("Character creation failed");
    }
  } catch (error) {
    console.error('Error creating character:', error);
  }
};

const updateCharacter = async () => {
  try {
    const id = route.params.id;
    console.log(character.value)
    const response = await fetchWithRedirect(`http://localhost:3000/api/characters/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(character.value),
    });
    if (response.ok) {
      router.push(`/characters/${id}`);
    } else {
      console.error("Character update failed");
    }
  } catch (error) {
    console.error('Error updating character:', error);
  }
};

const fetchCharacter = async () => {
  try {
    const id = route.params.id;
    const response = await fetchWithRedirect(`http://localhost:3000/api/characters/${id}`, {
      method: "GET",
    });
    if (response.ok) {
      const data = await response.json();
      character.value = data;
      selectedSystem.value = data.system;
    } else {
      console.error("Failed to fetch Character");
    }
  } catch (error) {
    console.error('Error fetching character:', error);
  }
}

if (editMode) {
  onMounted(fetchCharacter)
}
</script>
