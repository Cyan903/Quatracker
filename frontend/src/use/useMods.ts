const modIdentifiers = {
    None: BigInt("-1"),
    NoSliderVelocity: BigInt("1") << BigInt("0"),
    Speed05X: BigInt("1") << BigInt("1"),
    Speed06X: BigInt("1") << BigInt("2"),
    Speed07X: BigInt("1") << BigInt("3"),
    Speed08X: BigInt("1") << BigInt("4"),
    Speed09X: BigInt("1") << BigInt("5"),
    Speed11X: BigInt("1") << BigInt("6"),
    Speed12X: BigInt("1") << BigInt("7"),
    Speed13X: BigInt("1") << BigInt("8"),
    Speed14X: BigInt("1") << BigInt("9"),
    Speed15X: BigInt("1") << BigInt("10"),
    Speed16X: BigInt("1") << BigInt("11"),
    Speed17X: BigInt("1") << BigInt("12"),
    Speed18X: BigInt("1") << BigInt("13"),
    Speed19X: BigInt("1") << BigInt("14"),
    Speed20X: BigInt("1") << BigInt("15"),
    Strict: BigInt("1") << BigInt("16"),
    Chill: BigInt("1") << BigInt("17"),
    NoPause: BigInt("1") << BigInt("18"),
    Autoplay: BigInt("1") << BigInt("19"),
    Paused: BigInt("1") << BigInt("20"),
    NoFail: BigInt("1") << BigInt("21"),
    NoLongNotes: BigInt("1") << BigInt("22"),
    Randomize: BigInt("1") << BigInt("23"),
    Speed055X: BigInt("1") << BigInt("24"),
    Speed065X: BigInt("1") << BigInt("25"),
    Speed075X: BigInt("1") << BigInt("26"),
    Speed085X: BigInt("1") << BigInt("27"),
    Speed095X: BigInt("1") << BigInt("28"),
    Inverse: BigInt("1") << BigInt("29"),
    FullLN: BigInt("1") << BigInt("30"),
    Mirror: (BigInt("0") | (BigInt("1") << BigInt("31"))) >> BigInt("0"),
    Coop: BigInt("4294967296"),
    Speed105X: BigInt("8589934592"),
    Speed115X: BigInt("17179869184"),
    Speed125X: BigInt("34359738368"),
    Speed135X: BigInt("68719476736"),
    Speed145X: BigInt("137438953472"),
    Speed155X: BigInt("274877906944"),
    Speed165X: BigInt("549755813888"),
    Speed175X: BigInt("1099511627776"),
    Speed185X: BigInt("2199023255552"),
    Speed195X: BigInt("4398046511104"),
    HeatlthAdjust: BigInt("8796093022208"), // Test mod for making long note windows easier
    NoMiss: BigInt("17592186044416"),
};

const modText = {
    None: "",
    NoSliderVelocity: "NSV",
    Speed05X: "0.5x",
    Speed06X: "0.6x",
    Speed07X: "0.7x",
    Speed08X: "0.8x",
    Speed09X: "0.9x",
    Speed11X: "1.1x",
    Speed12X: "1.2x",
    Speed13X: "1.3x",
    Speed14X: "1.4x",
    Speed15X: "1.5x",
    Speed16X: "1.6x",
    Speed17X: "1.7x",
    Speed18X: "1.8x",
    Speed19X: "1.9x",
    Speed20X: "2.0x",
    Strict: "Strict",
    Chill: "Chill",
    NoPause: "No Pause",
    Autoplay: "AP",
    Paused: "Paused",
    NoFail: "NF",
    NoLongNotes: "NLN",
    Randomize: "RND",
    Speed055X: "0.55x",
    Speed065X: "0.65x",
    Speed075X: "0.75x",
    Speed085X: "0.85x",
    Speed095X: "0.95x",
    Inverse: "INV",
    FullLN: "FLN",
    Mirror: "MR",
    Coop: "Co-op",
    Speed105X: "1.05x",
    Speed115X: "1.15x",
    Speed125X: "1.25x",
    Speed135X: "1.35x",
    Speed145X: "1.45x",
    Speed155X: "1.55x",
    Speed165X: "1.65x",
    Speed175X: "1.75x",
    Speed185X: "1.85x",
    Speed195X: "1.95x",
    HeatlthAdjust: "LN Adjust",
    NoMiss: "No Miss",
};

// const m = 2694840385;
// 1679 test
export function useMods(m: string, text?: boolean) {
    const mods = [];

    try {
        if (parseInt(m) <= 0) {
            if (!text) mods.push("None");
            return mods;
        }

        for (const mod in modIdentifiers) {
            const idx = mod as keyof typeof modIdentifiers;

            if (modIdentifiers[idx] === BigInt("-1")) continue;
            if (BigInt(m) & modIdentifiers[idx]) {
                if (text) {
                    mods.push(modText[idx]);
                    continue;
                }

                mods.push(mod);
            }
        }
    } catch (e) {
        console.warn("[useMods] could not parse mods!", e);
        return ["Unknown"];
    }

    return mods;
}
