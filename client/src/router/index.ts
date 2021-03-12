import Vue from "vue";
import VueRouter from "vue-router";
import Generator from "../views/Generator.vue";
import Calendar from "@/views/Calendar.vue";

Vue.use(VueRouter);

export default new VueRouter({
  mode: "history",
  routes: [
    { path: "/", redirect: "/calendar" },
    { path: "/generator", component: Generator },
    { path: "/calendar", component: Calendar }
  ]
});
