export function useShorten(str: string, n: number) {
    if (str.length - 3 > n) {
        return str.slice(0, n) + "...";
    }

    return str;
}

export function useComma(n: string | number) {
    return n.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}
