<template>
  <div>
    <h1>{{ editMode ? 'Edit Skill' : 'Create New Skill' }}</h1>
    <form @submit.prevent="editMode ? updateSkill() : createSkill()">
      <p>Name: <input v-model="skill.name" placeholder="name" /></p>
      <p>Description: <input v-model="skill.description" placeholder="description" /></p>
      <button type="submit">{{ editMode ? 'Save Changes' : 'Create Skill' }}</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import fetchWithRedirect from '@/utils/fetchWithRedirect';
import { ref, onMounted } from 'vue';
import { useRouter, useRoute } from "vue-router";

const skill = ref<any>({});
const router = useRouter();
const route = useRoute();
const editMode = route.params.id;

const createSkill = async () => {
  try {
    const response = await fetchWithRedirect("http://localhost:3000/api/skills", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(skill.value),
    });
    if (response.ok) {
      router.push("/skills");
    } else {
      console.error("Skill creation failed");
    }
  } catch (error) {
    console.error('Error creating skill:', error);
  }
};

const updateSkill = async () => {
  try {
    const id = route.params.id;
    const response = await fetchWithRedirect(`http://localhost:3000/api/skills/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(skill.value),
    });
    if (response.ok) {
      router.push("/skills");
    } else {
      console.error("Skill update failed");
    }
  } catch (error) {
    console.error('Error updating skill:', error);
  }
};

const fetchSkill = async () => {
  try {
    const id = route.params.id;
    const response = await fetchWithRedirect(`http://localhost:3000/api/skills/${id}`, {
      method: "GET",
    });
    if (response.ok) {
      const data = await response.json();
      skill.value = data.skill;
    } else {
      console.error("Failed to fetch Skill");
    }
  } catch (error) {
    console.error('Error fetching skill:', error);
  }
}

if (editMode) {
  onMounted(fetchSkill)
}
</script>
