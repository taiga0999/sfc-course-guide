import Vue from 'vue';

const component = name => Vue.options.components[name];

export default () => [
  {
    path: '/',
    component: component('search'),
  },
  {
    path: '*',
    component: { template: '<div>404 Not Found</div>' },
  },
];
