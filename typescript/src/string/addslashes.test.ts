import { addslashes } from "./addslashes";

const textSingeQuote = "My name is O' Reilly";
const textDoubleQuote = "My name is O\" Reilly";

describe("Add slashes", () => {
    it("should add backslash before single quote", () => {
        expect(addslashes(textSingeQuote)).toEqual("My name is O\\' Reilly");
    });

    it("should add backslash before double quote", () => {
        expect(addslashes(textDoubleQuote)).toEqual("My name is O\" Reilly");
    });
});