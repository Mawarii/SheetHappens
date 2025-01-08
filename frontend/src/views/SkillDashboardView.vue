<template>
  <div class="dashboard">
    <h1>Skills <button class="add-btn" @click="createSkill" type="button"><OhVueIcon name="px-plus" scale="1.5" /></button></h1>
    <div v-if="skills" class="skill-list">
      <div v-for="skill in skills" :key="skill['id']" @click="goToSkillDetail(skill['id'])" class="skill-item">
        <span class="skill-name">{{ skill['name'] }}</span>
        <button class="delete-btn" @click.stop="confirmDelete(skill['id'])" type="button">
          <OhVueIcon name="bi-trash-fill" />
        </button>
      </div>
    </div>
    <div v-else>
      <span>No Skills available!</span>
    </div>
  </div>

  <div v-if="showModal" class="modal-overlay">
    <div class="modal">
      <h2>Are you sure you want to delete this skill?</h2>
      <div class="modal-actions">
        <button @click="deleteSkill" class="confirm-btn">Yes, delete</button>
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

const skills = ref<any | null>(null);
const showModal = ref(false);
const skillToDelete = ref<number | null>(null);
const router = useRouter();

const fetchSkills = async () => {
  try {
    const res = await fetchWithRedirect("http://localhost:3000/api/skills", {
      method: "GET",
    });
    const data = await res.json();
    skills.value = data.skills;
  } catch (error) {
    console.error('Error fetching skills:', error);
  }
};

const goToSkillDetail = (id: number) => {
  router.push(`/skills/${id}`);
};

const confirmDelete = (id: number) => {
  skillToDelete.value = id;
  showModal.value = true;
};

const cancelDelete = () => {
  showModal.value = false;
  skillToDelete.value = null;
};

const deleteSkill = async () => {
  try {
    if (skillToDelete.value !== null) {
      const res = await fetchWithRedirect(`http://localhost:3000/api/skills/${skillToDelete.value}`, {
        method: "DELETE",
      });
      if (res.ok) {
        showModal.value = false;
        skillToDelete.value = null;
        fetchSkills();
      }
    }
  } catch (error) {
    console.error('Error deleting skill:', error);
  }
};

const createSkill = () => {
  router.push("/skills/create");
}

onMounted(fetchSkills);
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

.skill-list {
  margin-top: 20px;
}

.skill-item {
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

.skill-item:hover {
  background-color: #252525;
}

.skill-item:has(.delete-btn:hover) {
  background-color: #000000;
}

.skill-name {
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
