import { createWebHistory, createRouter } from "vue-router";
import App from "../view/App.vue";
// import About from "@/views/About.vue";

const routes = [
  {
    path: "/",
    name: "App",
    component: App,
  },
  // {
  //   path: "/about",
  //   name: "About",
  //   component: About,
  // },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;