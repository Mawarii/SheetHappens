<template>
  <header>
      <nav>
        <RouterLink to="/">Login</RouterLink>
        <RouterLink to="/characters">Dashboard</RouterLink>
        <RouterLink to="/skills">Skills</RouterLink>
        <button @click="logout" class="logout-button">Logout</button>
      </nav>
  </header>

  <div class="content">
    <RouterView />
  </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView, useRouter } from 'vue-router'

const router = useRouter();

const logout = async () => {
  try {
    const response = await fetch("http://localhost:3000/api/auth/logout", {
      method: "GET",
      credentials: "include",
    });

    if (response.ok) {
      router.push("/");
    } else {
      console.error("Logout failed");
    }
  } catch (error) {
    console.error("Error during logout:", error);
  }
};
</script>

<style scoped>
header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  width: 100%;
  background: var(--color-background);
  z-index: 100;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 1rem 0;
}

nav {
  width: 100%;
  font-size: 14px;
  text-align: left;
}

nav a,
.logout-button {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
  background: none;
  color: var(--color-text);
  font: inherit;
  cursor: pointer;
  text-decoration: none;
}

nav a:first-of-type,
.logout-button:first-of-type {
  border: 0;
}

.logout-button:hover {
  text-decoration: underline;
}

/* Adjust the content to make room for the fixed header */
.content {
  margin-top: 60px;
  padding: 2rem;
}
</style>