import Vue from "vue";
import VueRouter from "vue-router";
import Generator from "../views/Generator.vue";

Vue.use(VueRouter);

export default new VueRouter({
  mode: "history",
  routes: [
    { path: "/", redirect: "/generator" },
    { path: "/generator", component: Generator }
  ]
});
