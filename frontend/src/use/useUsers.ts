import { GetUserJudgements } from "../../wailsjs/go/app/App";
import { useWails } from "./useWails";

export async function getJudgements() {
    return await useWails<string[]>(GetUserJudgements);
}
