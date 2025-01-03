<template>
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="SheetHappens" />

      <nav>
        <RouterLink to="/">Login</RouterLink>
        <button @click="logout" class="logout-button">Logout</button>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<script setup lang="ts">
import { RouterLink, RouterView, useRouter } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'

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
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
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

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;
    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
