import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/LoginView.vue'
import Register from '../views/RegisterView.vue'
import Dashboard from '../views/DashboardView.vue'
import CharacterDetail from '../views/CharacterDetailView.vue'
import CharacterForm from '../views/CharacterFormView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Login",
      component: Login,
    },
    {
      path: "/register",
      name: "Register",
      component: Register,
    },
    {
      path: "/characters",
      name: "Dashboard",
      component: Dashboard,
    },
    {
      path: "/characters/:id",
      name: "CharacterDetail",
      component: CharacterDetail,
    },
    {
      path: "/characters/:id/edit",
      name: "CharacterEdit",
      component: CharacterForm,
    },
    {
      path: "/characters/create",
      name: "CharacterCreate",
      component: CharacterForm,
    },
  ],
})

export default router
