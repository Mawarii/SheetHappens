<template>
  <div>
    <h1>{{ character['name'] }}
      <button class="edit-btn" @click.stop="goToEditMode()" type="button">
          <OhVueIcon name="fa-edit" />
      </button>
    </h1>
    <form v-if="character">
      <div v-for="(value, key) in character" :key="key">
        <div v-if="!['id', 'name', 'user_id', 'skills', 'craft'].includes(String(key)) && value">
          <span :class="key" class="key">{{ capitalizeFirstLetter(String(key)) }}: </span><span class="value">{{ value }}</span>
        </div>
        <div v-if="String(key) === 'skills'">
          {{ capitalizeFirstLetter(String(key)) }}:
          <div v-for="(skills, category) in value" :key="category">
            {{ category }}
            <div v-for="(skillValue, skill) in skills" :key="skill">
              <span>{{ skill }}: {{ skillValue }}</span>
            </div>
          </div>
        </div>
        <div v-if="String(key) === 'craft'">
          {{ capitalizeFirstLetter(String(key)) }}
          <div v-for="(craftValue, craft) in value" :key="craft">
            <span>{{ craft }}: {{ craftValue }}</span>
          </div>
        </div>
      </div>
    </form>
    <form v-else>
      Character not found!
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { OhVueIcon, addIcons } from "oh-vue-icons";
import { FaEdit } from "oh-vue-icons/icons";

addIcons(FaEdit)

const character = ref<any>({});
const route = useRoute();
const router = useRouter();

const fetchCharacter = async () => {
try {
    const id = route.params.id;
    const res = await fetch(`http://localhost:3000/api/characters/${id}`, {
        method: "GET",
        credentials: 'include',
      });
      const data = await res.json();
      character.value = data.character;
  } catch (error) {
    console.error('Error fetching character:', error);
  }
};

const capitalizeFirstLetter = (str: string) => {
  if (str == "mentalhealth") {
    return "Mental Health"
  }
  return str.charAt(0).toUpperCase() + str.slice(1);
};

const goToEditMode = () => {
    const id = route.params.id;
    router.push(`/characters/${id}/edit`)
}

onMounted(fetchCharacter);
</script>

<style scoped>
.edit-btn {
  background-color: rgb(0, 98, 255);
  color: white;
}

.edit-btn:hover {
  background-color: rgb(0, 73, 191);
}
</style>
