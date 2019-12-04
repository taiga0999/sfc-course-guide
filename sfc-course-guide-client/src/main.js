import Vue from 'vue';
import VueRouter from 'vue-router';
import axios from 'axios';
import setting from './config/setting.json';
import routes from './config/routes';
import Request from './plugins/requests';

// Globally import modules
const modules = [
  ['$axios', axios],
  ['$setting', setting],
];
modules.forEach(([prop, value]) => {
  Object.defineProperty(Vue.prototype, prop, { value });
});

/**
 * Recursively scan this directory for the Vue components and automatically
 * register them with their "basename".
 *
 * Eg. ./components/ExampleComponent.vue -> <example></example>
 */
const files = require.context('./', true, /\.vue$/i);
files.keys().forEach((key) => {
  Vue.component(
    key.replace(/(\.\/components\/)(.*)Component\.vue/, '$2') // Get the component name
      .replace(/([a-z])([A-Z])/g, '$1-$2') // Convert to kebab-case
      .replace(/([A-Z])([A-Z][a-z])/g, '$1-$2')
      .toLowerCase(), // Convert to lowercase
    files(key).default,
  );
});

// Use plugin
const plugins = [Request, VueRouter];
plugins.forEach(p => Vue.use(p));

const router = new VueRouter({
  mode: 'history',
  routes: routes(),
});

// Creating the Vue application instance
// eslint-disable-next-line no-unused-vars
const app = new Vue({
  router,
}).$mount('.vue');
