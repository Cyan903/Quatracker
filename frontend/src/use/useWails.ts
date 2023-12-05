import { LogError } from "@/../wailsjs/runtime/runtime";

export async function useWails<Type>(
    func: Function,
    ...args: any
): Promise<{
    error: unknown | null;
    result: Type | null;
}> {
    try {
        const result = await func(...args);

        return {
            error: null,
            result,
        };
    } catch (e: unknown) {
        LogError(`[useWails] wails called failed: ${e}`);

        return {
            error: e,
            result: null,
        };
    }
}
