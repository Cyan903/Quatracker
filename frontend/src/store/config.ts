import { defineStore } from "pinia";
import { computed, onMounted, ref } from "vue";
import { useToast } from "vue-toastification";

import { GetConfig, LoadDB, SetGamePath } from "../../wailsjs/go/app/App";
import { LogInfo } from "../../wailsjs/runtime/runtime";
import { config } from "../../wailsjs/go/models";
import { useWails } from "../use/useWails";

const toast = useToast();

export const useConfigStore = defineStore("config", () => {
    const err = ref(true);
    const data = ref<config.Config>({
        GamePath: "",
    });

    // TODO: Improve
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
    };
});
