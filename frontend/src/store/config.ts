import { defineStore } from "pinia";
import { computed, onMounted, ref } from "vue";
import { useToast } from "vue-toastification";

import { GetConfig, LoadDB, SetGamePath } from "../../wailsjs/go/app/App";
import { LogError, LogInfo } from "../../wailsjs/runtime/runtime";
import { config } from "../../wailsjs/go/models";

const toast = useToast();

export const useConfigStore = defineStore("config", () => {
    const data = ref<config.Config>();

    // TODO: Improve
    const validConfig = computed(() => {
        return data.value?.GamePath?.length != 0;
    });

    const connectSQL = async () => {
        if (!data.value?.GamePath) {
            LogInfo("Skipping connection test due to config being incomplete");
            return;
        }

        try {
            await LoadDB();
            LogInfo("Connected to quaver.db");
        } catch (e) {
            toast.error("Could not connect to quaver.db, is path correct?");

            LogError("[config.ts] could not connect to quaver.db");
            LogError(`[config.ts] ${e}`);
        }
    };

    const refresh = async () => {
        try {
            data.value = await GetConfig();

            LogInfo("Loaded configuration");
            await connectSQL();
        } catch (e) {
            toast.error("Could not load configuration file.");

            LogError("[config.ts] could not load config");
            LogError(`[config.ts] ${e}`);
        }
    };

    const setGamePath = async (path: string): Promise<boolean> => {
        try {
            await SetGamePath(path);
            await refresh();
            LogInfo(`Updated path -> ${path}`);

            return true;
        } catch (e) {
            LogError("[config.ts] could not set game path");
            LogError(`[config.ts] ${e}`);

            return false;
        }
    };

    onMounted(refresh);

    return {
        data,
        validConfig,
        refresh,
        setGamePath,
    };
});
