import placeholder from "@/assets/images/placeholder-bg.png";

export async function useImage(id: number) {
    const req = await fetch(`/id/${id}.id`)
        .then((r) => r.blob())
        .catch(() => console.info(`[BackgroundMap] ${id} failed`));

    if (!req) return placeholder;
    return URL.createObjectURL(req);
}
