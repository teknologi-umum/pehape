import { str_split } from "./str_split";

const str = "Hello World!";

describe("string split", () => {
  it("should be able to return every letter as element of the array", () => {
    expect(str_split(str)).toEqual(
        ["H", "e", "l", "l", "o", " ", "W", "o", "r", "l", "d", "!"]
    );
  });

  it("should be able to return 2 or more letter as element of the array depend on number of len", () => {
    expect(str_split(str, 2)).toEqual(["He", "ll", "o ", "Wo", "rl", "d!"]);
    expect(str_split(str, 3)).toEqual(["Hel", "lo ", "Wor", "ld!"]);
    expect(str_split(str, 4)).toEqual(["Hell", "o Wo", "rld!"]);
  });

  it("should be able to return false", () => {
    expect(str_split(str, 0)).toEqual(false);
    expect(str_split(str, -1)).toEqual(false);
  });

  it("should be able to return the entire string as the only element of the array", () => {
    expect(str_split(str, 15)).toEqual(["Hello World!"]);
  });
});
