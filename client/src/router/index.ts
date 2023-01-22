import { createRouter, createWebHistory } from "vue-router";
import GeneratorInput from "@/views/generator/GeneratorInput.vue";
import Generator from "@/views/generator/Generator.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: "/", redirect: "/generator" },
    { path: "/generator", component: GeneratorInput },
    { path: "/generator/:a_hours/:b_hours/:month/:year", component: Generator }
  ],
});

export default router;