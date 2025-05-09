import './main.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap';

import { createApp } from 'vue'
import App from './App.vue'
import router from './router/router';

createApp(App)
  .use(router)
  .mount('#app')
