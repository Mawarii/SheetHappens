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
          <h4>{{ category }}</h4>
          <div v-for="skill in skills" :key="skill['id']">
            <span>{{ getSkillNameById(skill['skill_id']) }}: </span>
            <input v-model.number="skill['value']" type="number" min="0" />
            <br />
            <!-- <input v-model.number="character.skills[category][skillid]" type="number" min="0" /> -->
          </div>
        </div>
        <h5>Category</h5>
        <p><input v-model="skillCategory" placeholder="category" /></p>
        <select v-model="selectedSkill">
          <option disabled value="">Please select one</option>
          <option v-for="skill in allSkills" :key="skill['id']" :value="skill['id']">
            {{ skill['name'] }}
          </option>
        </select>
        <p><input v-model="skillValue" placeholder="value" /></p>
        <button type="button" @click="addSkill">Add Skill</button>
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

const character = ref<any>({
  skills: {}
});
const allSkills = ref<any>({});
const skillCategory = ref("");
const skillValue = ref(0);
const selectedSkill = ref("");
const router = useRouter();
const route = useRoute();
const editMode = route.params.id;

const getSkillNameById = (skillId: string) => {
  const skill = allSkills.value.find((s: { id: string }) => s.id === skillId);
  return skill ? skill.name : "Unknown Skill";
};

const addSkill = () => {
  if (!character.value.skills) {
    character.value.skills = {};
  }
  if (skillCategory.value && selectedSkill.value) {
    if (!character.value.skills[skillCategory.value]) {
      character.value.skills[skillCategory.value] = [];
    }

    character.value.skills[skillCategory.value].push({
      skill_id: selectedSkill.value,
      value: skillValue.value
    });

    selectedSkill.value = '';
    skillValue.value = 0;
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
      // const id = data.result["InsertedID"]
      // router.push(`/characters/${id}`);
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
      character.value = data.character;
    } else {
      console.error("Failed to fetch Character");
    }
  } catch (error) {
    console.error('Error fetching character:', error);
  }
}

const fetchSkills = async () => {
  try {
    const res = await fetchWithRedirect("http://localhost:3000/api/skills", {
      method: "GET",
    });
    const data = await res.json();
    allSkills.value = data.skills;
  } catch (error) {
    console.error('Error fetching skills:', error);
  }
};

if (editMode) {
  onMounted(fetchCharacter)
}
onMounted(fetchSkills)
</script>
