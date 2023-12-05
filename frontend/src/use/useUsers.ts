import { GetUserJudgements } from "@/../wailsjs/go/app/App";
import { useWails } from "@/use/useWails";

export async function getJudgements() {
    return await useWails<string[]>(GetUserJudgements);
}
