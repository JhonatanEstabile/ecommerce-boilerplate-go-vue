import { createWebHistory, createRouter } from 'vue-router';

const routes = [
  { path: '/login', component: () => import('../views/LoginView.vue')},
  { path: '/admin/products', component: () => import('../views/ProductsView.vue')},
  { path: '/', redirect: '/login' },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
