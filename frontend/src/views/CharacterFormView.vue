<template>
  <div>
    <h1>{{ editMode ? 'Edit Character' : 'Create New Character' }}</h1>
    <form @submit.prevent="editMode ? updateCharacter() : createCharacter()">
      <p>Name: <input v-model="character.name" placeholder="name" /></p>
      <p>Level: <input v-model="character.level" placeholder="level" type="number" /></p>
      <p>Health: <input v-model="character.health" placeholder="health" type="number" /></p>
      <p>Mental Health: <input v-model="character.mentalhealth" placeholder="mental health" type="number"/></p>
      <p>Mana: <input v-model="character.mana" placeholder="mana" type="number"/></p>
      <p>Race: <input v-model="character.race" placeholder="race" /></p>
      <p>Gender: <input v-model="character.gender" placeholder="gender" /></p>
      <p>Height: <input v-model="character.height" placeholder="height" /></p>
      <p>Weight: <input v-model="character.weight" placeholder="weight" /></p>
      <p>Dodge: <input v-model="character.dodge" placeholder="dodge" type="number" /></p>
      <div>
        <h3>Skills</h3>
        <div v-for="(skills, category) in character.skills" :key="category">
          {{ category }}
          <div v-for="(_, skill) in skills" :key="skill">
            <span>{{ skill }}: </span>
            <input v-model.number="character.skills[category][skill]" type="number" min="0" />
            <button type="button" @click="removeSkill(String(category), String(skill))">Remove</button>
          </div>
        </div>
        <div>
          <input v-model="newSkillCategory" placeholder="Category" />
          <input v-model="newSkillName" placeholder="Skill" />
          <input v-model.number="newSkillValue" type="number" placeholder="Value" min="0" />
          <button type="button" @click="addSkill()">Add Skill</button>
        </div>
      </div>
      <div>
        <h3>Craft</h3>
        <div v-for="(_, craft) in character.craft" :key="craft">
          {{ craft }}: 
            <input v-model.number="character.craft[craft]" type="number" min="0" />
            <button type="button" @click="removeCraft(String(craft))">Remove</button>
        </div>
        <div>
          <input v-model="newCraftName" placeholder="Craft" />
          <input v-model.number="newCraftValue" type="number" placeholder="Value" min="0" />
          <button type="button" @click="addCraft()">Add Craft</button>
        </div>
      </div>
      <button type="submit">{{ editMode ? 'Save Changes' : 'Create Character' }}</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import fetchWithRedirect from '@/utils/fetchWithRedirect';
import { ref, onMounted } from 'vue';
import { useRouter, useRoute } from "vue-router";

const character = ref<any>({
  skills: {},
  craft: {}
});

const newSkillCategory = ref('');
const newSkillName = ref('');
const newSkillValue = ref(0);
const newCraftName = ref('');
const newCraftValue = ref(0);
const router = useRouter();
const route = useRoute();
const editMode = route.params.id;

const addSkill = () => {
  if (newSkillCategory.value && newSkillName.value) {
    if (!character.value.skills[newSkillCategory.value]) {
      character.value.skills[newSkillCategory.value] = {};
    }
    character.value.skills[newSkillCategory.value][newSkillName.value] = newSkillValue.value;
    newSkillName.value = '';
    newSkillValue.value = 0;
  }
};

const removeSkill = (category: string, skill: string) => {
  if (character.value.skills[category]) {
    delete character.value.skills[category][skill];
    if (Object.keys(character.value.skills[category]).length === 0) {
      delete character.value.skills[category];
    }
  }
};

const addCraft = () => {
  if (!character.value.craft) {
    character.value.craft = {};
  }
  if (newCraftName.value) {
    character.value.craft[newCraftName.value] = newCraftValue.value;
    newCraftName.value = '';
    newCraftValue.value = 0;
  }
};

const removeCraft = (craft: string) => {
  if (character.value.craft[craft]) {
    delete character.value.craft[craft];
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
      const id = data.result["InsertedID"]
      router.push(`/characters/${id}`);
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
      character.value = data.character;
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
