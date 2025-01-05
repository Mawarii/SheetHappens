<template>
  <div>
    <h1>{{ editMode ? 'Edit Character' : 'Create New Character' }}</h1>
    <form @submit.prevent="editMode ? updateCharacter() : createCharacter()">
      <p><strong>Name:</strong> <input v-model="character.name" placeholder="name" /></p>
      <p><strong>Level:</strong> <input v-model="character.level" placeholder="level" type="number" /></p>
      <p><strong>Health:</strong> <input v-model="character.health" placeholder="health" type="number" /></p>
      <p><strong>Mental Health:</strong> <input v-model="character.mentalhealth" placeholder="mental health" type="number"/></p>
      <p><strong>Mana:</strong> <input v-model="character.mana" placeholder="mana" type="number"/></p>
      <p><strong>Race:</strong> <input v-model="character.race" placeholder="race" /></p>
      <p><strong>Gender:</strong> <input v-model="character.gender" placeholder="gender" /></p>
      <p><strong>Height:</strong> <input v-model="character.height" placeholder="height" /></p>
      <p><strong>Weight:</strong> <input v-model="character.weight" placeholder="weight" /></p>
      <p><strong>Dodge:</strong> <input v-model="character.dodge" placeholder="dodge" type="number" /></p>
      <div>
        <h3>Skills</h3>
        <div v-for="(skills, category) in character.skills" :key="category">
          <strong>{{ category }}</strong>
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
          <button type="button" @click="addSkill">Add Skill</button>
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
  skills: {}
});

const newSkillCategory = ref('');
const newSkillName = ref('');
const newSkillValue = ref(0);
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
