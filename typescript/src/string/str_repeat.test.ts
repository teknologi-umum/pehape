import { str_repeat } from "./str_repeat";
const str = "Hello";

describe("string repeat", () => {
  it("have string double", () => {
    expect(str_repeat(str, 2)).toEqual("HelloHello");
  });
  it("empty string", () => {
    expect(str_repeat(str, 0)).toEqual("");
  });
  it("invalid second parameter", () => {
    expect(() => str_repeat(str, "a")).toThrow("Argument is must be number");
  });
});
