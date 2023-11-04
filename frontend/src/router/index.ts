import { createRouter, createWebHistory } from "vue-router";

import ProfileView from "../views/ProfileView.vue";
import AboutView from "../views/AboutView.vue";
import NotFoundView from "../views/ErrorView.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "profile",
            component: ProfileView,
            meta: { title: "Profile" },
        },

        {
            path: "/about",
            name: "about",
            component: AboutView,
            meta: { title: "About" },
        },

        {
            path: "/:pathMatch(.*)*",
            name: "not-found",
            component: NotFoundView,
            meta: { title: "Not Found" },
        },
    ],
});

router.beforeEach((to) => {
    document.title = `${to.meta.title || "Unknown"} | QuaverBuddy`;
});

export default router;
