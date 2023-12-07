import { defineStore } from "pinia";
import { computed, onMounted, ref } from "vue";
import { useToast } from "vue-toastification";

import { useWails } from "@/use/useWails";
import { LogInfo } from "@wails/runtime/runtime";
import { config } from "@wails/go/models";
import { GetConfig, LoadDB, SetGamePath, SetMainMode } from "@wails/go/app/App";

const toast = useToast();

export const useConfigStore = defineStore("config", () => {
    const err = ref(true);
    const data = ref<config.Config>({
        GamePath: "",
        MainMode: false,
    });

    const validConfig = computed(() => {
        return !err.value;
    });

    const setGamePath = async (path: string): Promise<boolean> => {
        const game = await useWails(SetGamePath, path);

        if (game.error) {
            err.value = true;
            toast.error("Could not set game path!");

            return true;
        }

        await refresh();
        return err.value;
    };

    const setMainMode = async (mode: boolean): Promise<boolean> => {
        const main = await useWails(SetMainMode, mode);

        if (main.error) {
            err.value = true;
            toast.error("Could not set main mode!");

            return true;
        }

        await refresh();
        return err.value;
    };

    const refresh = async () => {
        const cfgs = await useWails<config.Config>(GetConfig);

        if (cfgs.error || !cfgs.result) {
            err.value = true;

            if (cfgs.error) {
                toast.error("Could not import configuration!");
            }

            return;
        }

        data.value = cfgs.result;

        // Test Connection
        const conn = await useWails(LoadDB);

        if (conn.error) {
            err.value = true;
            toast.error(
                "Could not connect to database, is the game path correct?",
            );

            return;
        }

        LogInfo("[config.ts] loaded configuration");
        err.value = false;
    };

    onMounted(refresh);

    return {
        data,
        validConfig,

        refresh,
        setGamePath,
        setMainMode,
    };
});
