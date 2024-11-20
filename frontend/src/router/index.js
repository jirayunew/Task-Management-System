import { createRouter, createWebHistory } from "vue-router";
import AppView from "../views/AppView.vue";
import LoginView from "../views/Login.vue";

const routes = [
  {
    path: "/login",
    name: "Login",
    component: LoginView, // หน้า Login
    meta: { requiresAuth: false },
  },
  {
    path: "/",
    name: "App",
    component: AppView, // หน้า Task Board
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Middleware ตรวจสอบ JWT
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token");
  if (to.meta.requiresAuth && !token) {
    next("/login"); // ถ้าไม่มี JWT ให้ Redirect ไปหน้า Login
  } else if (to.path === "/login" && token) {
    next("/"); // ถ้ามี JWT แล้ว แต่พยายามเข้า /login ให้ Redirect ไป /
  } else {
    next(); // อนุญาตให้ผ่าน
  }
});

export default router;
