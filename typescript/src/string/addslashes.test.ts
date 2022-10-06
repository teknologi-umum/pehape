import { addslashes } from "./addslashes";

const textSingeQuote = "My name is O' Reilly";
const textDoubleQuote = "My name is O\" Reilly";

describe("Replace slashes", () => {
    it("should replace single quote text to backslash", () => {
        expect(addslashes(textSingeQuote)).toEqual("My name is O\\' Reilly");
    });

    it("should replace double quote text to backslash", () => {
        expect(addslashes(textDoubleQuote)).toEqual("My name is O\\' Reilly");
    });
});