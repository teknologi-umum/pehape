import { rtrim } from "./rtrim";

describe("rtrim with default characters", () => {
  it("should return trimmed string", () => {
    expect(rtrim("Hello, World!    ")).toStrictEqual("Hello, World!");
    expect(rtrim("Hello, World!\n  \t ")).toStrictEqual("Hello, World!");
    expect(rtrim("Hello, World!\0 \v")).toStrictEqual("Hello, World!");
  });

  it("should not trim the left side of the string", () => {
    expect(rtrim("    Hello, World!")).toStrictEqual("    Hello, World!");
    expect(rtrim("    Hello, World!    ")).toStrictEqual("    Hello, World!");
  });

  it("should return the string if there is nothing to be trimmed", () => {
    expect(rtrim("Hello, World!")).toStrictEqual("Hello, World!");
  });
});

describe("rtrim with different characters", () => {
  it("should return trimmed string with different characters", () => {
    expect(rtrim("Hello, World!rrrrrrr", "r")).toStrictEqual("Hello, World!");
    expect(rtrim("Hello, World!rtrtrttrr", "rt")).toStrictEqual("Hello, World!");
    expect(rtrim("Hello, World!0293174856", "0123456789")).toStrictEqual("Hello, World!");
  });

  it("should not trim the left side of the string", () => {
    expect(rtrim("rrrrrrHello, World!", "r")).toStrictEqual("rrrrrrHello, World!");
    expect(rtrim("rrrrrHello, World!rrrrr", "r")).toStrictEqual("rrrrrHello, World!");
  });

  it("should trim regex special characters", () => {
    expect(rtrim("Hello, World!+.*", "*.+")).toStrictEqual("Hello, World!");
    expect(rtrim("Hello, World![*]$", "*$[]")).toStrictEqual("Hello, World!");
  });
});

describe("rtrim with character range", () => {
  it("should return trimmed string for ranged characters", () => {
    expect(rtrim("\x09Hello, World!\x0A", "\x00..\x1F")).toStrictEqual("\x09Hello, World!");
    expect(rtrim("Hello, World!0293174856", "0..9")).toStrictEqual("Hello, World!");
  });

  it("should also work with reverse range", () => {
    expect(rtrim("Hello, World!aiaxmioxj", "z..a")).toStrictEqual("Hello, World!");    
  });

  it("should also work with multiple ranges", () => {
    expect(rtrim("Hello, World!38", "1..57..9")).toStrictEqual("Hello, World!");
    expect(rtrim("Hello, World!6", "1..57..9")).toStrictEqual("Hello, World!6");
  });
});