import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/start/:id',
    name: 'Start',
    component: () => import(/* webpackChunkName: "about" */ '../views/Start.vue'),
    children: [
      {
        path: 'loading',
        name: 'Load',
        component: () => import(/* webpackChunkName: "about" */ '../views/Start/Load.vue'),
        props: (route) => ({
          ...route.params
        })
      },
      {
        path: 'question',
        name: 'Question',
        component: () => import(/* webpackChunkName: "about" */ '../views/Start/Question.vue'),
        props: (route) => ({
          ...route.params
        })
      },
      {
        path: 'score',
        name: 'Score',
        component: () => import(/* webpackChunkName: "about" */ '../views/Start/Score.vue'),
        props: (route) => ({
          ...route.params
        })
      },
      {
        path: '',
        redirect: 'loading'
      }
    ]
  },
  {
    path: '/new',
    name: 'New',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/New.vue')
  }
];

const router = new VueRouter({
  mode: 'history',
  // base: process.env.BASE_URL+"#/",
  routes
});

export default router
