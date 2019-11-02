import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import MainCategory from "../views/MainCategory.vue";
import SubCategory from "../views/SubCategory.vue";
import Content from "../views/Content.vue";
import NoContent from "../views/NoContent.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home
  },
  {
    path: "/category/:mainID",
    name: "mainCategory",
    component: MainCategory
  },
  { path: "/category/sub/:subID", name: "subCategory", component: SubCategory },
  {
    path: "/content/:contentID",
    name: "content",
    component: Content
  },
  {
    path: "*",
    name: "noContent",
    component: NoContent
  }
];

const router = new VueRouter({
  routes
});

export default router;
