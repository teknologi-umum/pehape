import { explode } from "./explode";

const text = "Lazy brown fox jumps";

describe("no limit is specified", () => {
  it("should return explode string into array", () => {
    expect(explode(" ", text)).toEqual(["Lazy", "brown", "fox", "jumps"]);
    expect(explode("fox", text)).toEqual(["Lazy brown ", " jumps"]);
  });
});

describe("limit is specified", () => {
  it("should limit output to 'limit' elements when it's positive integer", () => {
    expect(explode(" ", text, 5)).toEqual(["Lazy", "brown", "fox", "jumps"]);
    expect(explode(" ", text, 3)).toEqual(["Lazy", "brown", "fox"]);
  });

  it("should exclude last -'limit' elements when it's negative integer", () => {
    expect(explode(" ", text, -1)).toEqual(["Lazy", "brown", "fox"]);
    expect(explode(" ", text, -3)).toEqual(["Lazy"]);
    expect(explode(" ", text, -5)).toEqual([]);
  });

  it("should return first element when 'limit' is 0", () => {
    expect(explode(" ", text, 0)).toEqual(["Lazy"]);
  });
});

describe("whitespaces", () => {
  it("should return empty string when exploding double space with single space separator", () => {
    expect(explode(" ", "Lazy  brown  fox")).toEqual([
      "Lazy",
      "",
      "brown",
      "",
      "fox",
    ]);
  });

  it("should not explode single space when separator is double space", () => {
    expect(explode("  ", "Lazy brown fox")).toEqual(["Lazy brown fox"]);
  });
});

describe("invalid parameters", () => {
  it("should throw error if separator is empty string", () => {
    expect(() => explode("", "any string")).toThrow(Error);
    expect(() => explode("", "any string")).toThrow(
      "Argument limit must not be empty string"
    );
  });
});
