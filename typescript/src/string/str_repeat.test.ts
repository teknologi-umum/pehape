import { str_repeat } from "./str_repeat";
const str = "Hello";

describe("string repeat", () => {
  it("should be able to repeat the string twice", () => {
    expect(str_repeat(str, 2)).toEqual("HelloHello");
  });
  it("should be able to ignore empty string", () => {
    expect(str_repeat(str, 0)).toEqual("");
  });
  it("should throw when invalid parameter is passed", () => {
    expect(() => {
      // @ts-expect-error because the second parameter should be a number but we pass a string for testing purpose
      str_repeat(str, "a");
    }).toThrow("Argument must be a number");
  });
});
