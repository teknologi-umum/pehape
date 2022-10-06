/**
 * Append backslashes to single quote, double quote
 * @param {string} text
 * @returns {string}
 */
export const addslashes = (text: string): string => {
    const replacedCondition: string[] = ["'", "\"\""];

    const temp: string[] = [];

    for (let i = 0; i < text.length; i++) {
        const replaced: boolean = replacedCondition.some((element) => text[i].includes(element));

        if (replaced) {
            temp.push("\\");
        }

        temp.push(text[i]);
    }

    return temp.join("");
};