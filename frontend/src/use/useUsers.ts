import { GetUserJudgements, GetUsers } from "@wails/go/app/App";
import { useWails } from "@/use/useWails";

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

export async function getUsers() {
    return useWails<UsersList>(GetUsers);
}

export async function getJudgements() {
    return await useWails<string[]>(GetUserJudgements);
}
