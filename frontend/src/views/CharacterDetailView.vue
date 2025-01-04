<template>
  <div>
    <h1>Character Details
      <label class="switch"><input type="checkbox" v-model="edit" ><span class="slider round"></span></label>
    </h1>
    <form v-if="character && !edit">
      <p><strong>Name:</strong> {{ character['name'] }}</p>
      <p v-if="character['level']"><strong>Level:</strong> {{ character['level'] }}</p>
      <p v-if="character['health']"><strong>Health:</strong> {{ character['health'] }}</p>
      <p v-if="character['mentalhealth']"><strong>Mental Health:</strong> {{ character['mentalhealth'] }}</p>
      <p v-if="character['race']"><strong>Race:</strong> {{ character['race'] }}</p>
      <p v-if="character['gender']"><strong>Gender:</strong> {{ character['gender'] }}</p>
      <p v-if="character['height']"><strong>Height:</strong> {{ character['height'] }}</p>
      <p v-if="character['weight']"><strong>Weight:</strong> {{ character['weight'] }}</p>
      <p v-if="character['dodge']"><strong>Dodge:</strong> {{ character['dodge'] }}</p>
    </form>
    <form v-else-if="edit" @submit.prevent="updateCharacter">
      <p><strong>Name:</strong> <input v-model="character.name" placeholder="name" /></p>
      <p><strong>Level:</strong> <input v-model="character.level" placeholder="level" type="number" /></p>
      <p><strong>Health:</strong> <input v-model="character.health" placeholder="health" type="number" /></p>
      <p><strong>Mental Health:</strong> <input v-model="character.mentalhealth" placeholder="mental health" type="number"/></p>
      <p><strong>Race:</strong> <input v-model="character.race" placeholder="race" /></p>
      <p><strong>Gender:</strong> <input v-model="character.gender" placeholder="gender" /></p>
      <p><strong>Height:</strong> <input v-model="character.height" placeholder="height" /></p>
      <p><strong>Weight:</strong> <input v-model="character.weight" placeholder="weight" /></p>
      <p><strong>Dodge:</strong> <input v-model="character.dodge" placeholder="dodge" type="number" /></p>
      <button type="submit">Save</button>
    </form>
    <form v-else>
      Something went wrong!
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';

const edit = ref<boolean>(false);
const character = ref<any | null>(null);

const route = useRoute();

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

const updateCharacter = async () => {
try {
    const id = route.params.id;
    const res = await fetch(`http://localhost:3000/api/characters/${id}`, {
        method: "PUT",
        credentials: 'include',
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: character.value.name,
          level: Number(character.value.level),
          health: Number(character.value.health),
          mentalhealth: Number(character.value.mentalhealth),
          race: character.value.race,
          gender: character.value.gender,
          height: character.value.height,
          weight: character.value.weight,
          dodge: Number(character.value.dodge),
        }),
      });
      if (res.ok) {
        edit.value = false;
      }
  } catch (error) {
    console.error('Error fetching character:', error);
  }
};

onMounted(fetchCharacter);
</script>

<style scoped>
/* The switch - the box around the slider */
.switch {
  position: relative;
  display: inline-block;
  width: 25px; /* Halbe Größe */
  height: 14px; /* Halbe Größe */
  vertical-align: middle; /* Vertikal mittig */
}

/* Hide default HTML checkbox */
.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

/* The slider */
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  border-radius: 34px; /* Sicherstellen, dass der Slider rund bleibt */
  transition: 0.2s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 12px; /* Reduziert für kleinere Größe */
  width: 12px; /* Reduziert für kleinere Größe */
  left: 1px; /* Positioniert den Schieberegler */
  bottom: 1px; /* Positioniert den Schieberegler */
  background-color: white;
  border-radius: 50%; /* Rundet den Schieberegler */
  transition: 0.2s;
}

input:checked + .slider {
  background-color: #2196F3;
}

input:focus + .slider {
  box-shadow: 0 0 1px #2196F3;
}

input:checked + .slider:before {
  transform: translateX(11px); /* Die Bewegung wurde angepasst */
}

/* Rounded sliders */
.slider.round {
  border-radius: 50px;
}

.slider.round:before {
  border-radius: 50%;
}
</style>
