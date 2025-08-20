export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  css: ["~/assets/css/pico.violet.min.css"],
  runtimeConfig: {
    public: {
      apiUrl: "http://localhost:8080/api",
    },
  },
})
