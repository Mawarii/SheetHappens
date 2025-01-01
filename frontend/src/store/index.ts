import { defineStore } from "pinia";

export const useTokenStore = defineStore("token", {
  state: () => ({
    token: null,
  }),

  actions: {
    async saveToken(token: null) {
      this.token = token;
      return this.token;
    },
  },
});
