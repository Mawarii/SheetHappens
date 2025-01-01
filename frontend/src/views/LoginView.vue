<template>
  <div>
    <h1>LOGIN</h1>
    <form @submit.prevent="login">
      <input v-model="username" placeholder="username" />
      <br />
      <br />
      <input v-model="password" placeholder="password" type="password" />
      <br />
      <br />
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
import { useTokenStore } from "../store/index";
export default {
  setup() {
    const tokenStore = useTokenStore();
    return { tokenStore };
  },
  data: () => {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    async login(e) {
      e.preventDefault();
      const response = await fetch("http://localhost:3000/api/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: this.username,
          password: this.password,
        }),
      });
      const { token } = await response.json();
      this.tokenStore.saveToken(token);
      this.$router.push("/characters");
    },
  },
};
</script>
