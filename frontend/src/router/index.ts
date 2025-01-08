import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/LoginView.vue'
import Register from '../views/RegisterView.vue'
import CharacterDashboard from '../views/CharacterDashboardView.vue'
import CharacterDetail from '../views/CharacterDetailView.vue'
import CharacterForm from '../views/CharacterFormView.vue'
import SkillDashboard from '@/views/SkillDashboardView.vue'
import SkillDetail from '@/views/SkillDetailView.vue'
import SkillForm from '@/views/SkillFormView.vue'

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
      component: CharacterDashboard,
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
    {
      path: "/skills",
      name: "SkillDashboard",
      component: SkillDashboard,
    },
    {
      path: "/skills/:id",
      name: "SkillDetail",
      component: SkillDetail,
    },
    {
      path: "/skills/:id/edit",
      name: "SkillEdit",
      component: SkillForm,
    },
    {
      path: "/skills/create",
      name: "SkillCreate",
      component: SkillForm,
    },
  ],
})

export default router
