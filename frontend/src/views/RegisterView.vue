<template>
  <div>
    <h1>Register</h1>
    <form @submit.prevent="register">
      <input v-model="username" placeholder="username" />
      <br />
      <br />
      <input v-model="password" placeholder="password" type="password" />
      <br />
      <br />
      <button type="submit">Register</button>
      <div v-if="success || errorMessage">
        <span v-if="success" class="success">{{ success }}</span>
        <span v-else class="error">{{ errorMessage }}</span>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const username = ref<string>("");
const password = ref<string>("");
const success = ref<string>("");
const errorMessage = ref<string>("");

const register = async () => {
  try {
    const response = await fetch("http://localhost:3000/api/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({
        username: username.value,
        password: password.value,
      }),
    });

    if (response.ok) {
      success.value = "Registered successfully";
    } else {
      const errorData = await response.json();
      errorMessage.value = errorData.message || "Failed to register";
    }
  } catch (error) {
    console.error("Error during registration:", error);
  }
};
</script>

<style scoped>
.success {
  color: green;
}

.error {
  color: red;
}
</style>
