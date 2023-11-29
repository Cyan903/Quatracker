import { createRouter, createWebHistory } from "vue-router";

import ProfileView from "../views/ProfileView.vue";
import ConfigView from "../views/ConfigView.vue";
import NotFoundView from "../views/404View.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "Profile",
            component: ProfileView,
            meta: { title: "Profile" },
        },

        {
            path: "/config",
            name: "config",
            component: ConfigView,
            meta: { title: "Configuration" },
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
