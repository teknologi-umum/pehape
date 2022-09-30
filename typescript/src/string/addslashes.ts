/* eslint-disable no-control-regex */
export const addslashes = (str: string): string => {
    return str.replace(/[\\"']/g, "\\$&").replace(/\u0000/g, "\\0");
};