import Vue from "vue";
import VueRouter from "vue-router";
import Generator from "../views/Generator/Generator.vue";
import GeneratorInput from "../views/Generator/GeneratorInput.vue";

Vue.use(VueRouter);

export default new VueRouter({
  mode: "history",
  routes: [
    { path: "/", redirect: "/generator" },
    { path: "/generator", component: GeneratorInput },
    { path: "/generator/:hours/:month/:year", component: Generator }
  ]
});
