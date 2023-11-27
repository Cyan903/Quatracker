import { reactive } from "vue";
import { useToast } from "vue-toastification";

import { GetUsers } from "../../wailsjs/go/app/App";
import { useWails } from "../use/useWails";

const toast = useToast();

export const user = reactive({
    id: -1,
    username: "",
});

export interface UsersList {
    Local: {
        Id: string;
        Username: string;
    }[];

    Unknown: {
        Id: string;
        Username: string;
    }[];
}

export async function getUsers(): Promise<UsersList> {
    const data = await useWails<UsersList>(GetUsers);
    const none = { Local: [], Unknown: [] };

    if (data.error) {
        toast.error("Could not get users!");
        return none;
    }

    return data.result || none;
}

export function setUser(id: number, name: string) {
    user.id = id;
    user.username = name;
}

export async function initUsers() {
    setUser(-1, "");

    const users = await getUsers();

    if (!users.Local?.length && !users.Unknown?.length) {
        toast.error("No users found!");
        return;
    }

    if (users.Local?.length !== 0) {
        setUser(parseInt(users.Local[0].Id), users.Local[0].Username);
        return;
    }

    setUser(parseInt(users.Unknown[0].Id), users.Unknown[0].Username);
};
