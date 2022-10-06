/**
 * Append backslashes to single quote, double quote
 * @param {string} text
 * @returns {string}
 */
export const addslashes = (text: string): string => {
    const appendCondition: string[] = ["'", "\"\""];

    const temp: string[] = [];

    for (let i = 0; i < text.length; i++) {
        const needToAppend: boolean = appendCondition.some((element) => text[i].includes(element));

        if (needToAppend) {
            temp.push("\\");
        }

        temp.push(text[i]);
    }

    return temp.join("");
};