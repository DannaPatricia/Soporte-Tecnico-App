import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import FormSupportCaseView from '@/views/FormSupportCaseView.vue'
import ResponseSupportCaseView from '../views/ResponseSupportCaseView.vue'
import ErrorView from '@/views/ErrorView.vue'
import CaseListView from '@/views/CaseListView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/report-list-view/:id',
      name: 'report-list-view',
      component: CaseListView,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/form-support-case',
      name: 'form-support-case',
      component: FormSupportCaseView,
    },
    {
      path: '/response-support-case/:id',
      name: 'response-support-case',
      component: ResponseSupportCaseView,
    },
    {
      path: '/:catchAll(.*)',
      name: 'ErrorView',
      component: ErrorView
    },
  ],
})

export default router
