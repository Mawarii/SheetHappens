import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/LoginView.vue'
import Dashboard from '../views/DashboardView.vue'
import CharacterDetail from '../views/CharacterDetailView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Login",
      component: Login,
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
  ],
})

export default router
