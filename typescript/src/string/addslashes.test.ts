/* eslint-disable quotes */
/* eslint-disable no-useless-escape */
import { addslashes } from "./addslashes";

describe("empty parameters", () => {
  it("should return empty string", () => {
    expect(addslashes("")).toBe("");
  });
});

describe("string without any characters", () => {
    it("should return string without any characters", () => {
        expect(addslashes("Hello world")).toBe("Hello world");
    });
});

describe("string with single quotes", () => {
  it("should return string with double backslashed", () => {
    expect(addslashes("'")).toBe("\\'");
    expect(addslashes("how're you doing?")).toBe("how\\'re you doing?");
    expect(addslashes("don't disturb u'r neighbours")).toBe("don\\'t disturb u\\'r neighbours");
    expect(addslashes("don't disturb u'r neighbours''")).toBe("don\\'t disturb u\\'r neighbours\\'\\'");
    expect(addslashes("\'")).toBe("\\'");
    expect(addslashes("'")).toBe("\\'");
  });
});

describe("string with double quotes", () => {
    it("should return string with double backslashed", () => {
        expect(addslashes(`he said, "he will be on leave"`)).toBe("he said, \\\"he will be on leave\\\"");
        expect(addslashes(`he said, ""he will be on leave"`)).toBe("he said, \\\"\\\"he will be on leave\\\"");
        expect(addslashes(`"""pehape"""`)).toBe("\\\"\\\"\\\"pehape\\\"\\\"\\\"");
        expect(addslashes(`"`)).toBe("\\\"");
        expect(addslashes(`hello\"`)).toBe("hello\\\"");
    });
});

describe("string with hexadecimal number", () => {
    it("should return string with double backslashed", () => {
        expect(addslashes(`0xABCDEF0123456789`)).toBe("0xABCDEF0123456789");
        expect(addslashes(`\xabcdef0123456789`)).toBe("«cdef0123456789");
        expect(addslashes(`!@#$%&*@$%#&/;:,<>`)).toBe("!@#$%&*@$%#&/;:,<>");
        expect(addslashes(`hello\x00world`)).toBe("hello\\0world");
    });
});