import { createMemoryHistory, createRouter } from 'vue-router';

const routes = [
  { path: '/', component: () => import('../views/HomeView.vue')},
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

export default router;
