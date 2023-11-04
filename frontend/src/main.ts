import "vue-toastification/dist/index.css";
import "./assets/styles/main.css";

import { createApp } from "vue";
import Toast, { POSITION } from "vue-toastification";

import App from "./App.vue";
import router from "./router";

const app = createApp(App);

app.use(Toast, {
    position: POSITION.BOTTOM_LEFT,
    timeout: 1500,
});

app.use(router);
app.mount("#app");
