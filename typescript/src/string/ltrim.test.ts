import { ltrim } from "./ltrim";

describe("ltrim with default characters", () => {
  it("should return trimmed string", () => {
    expect(ltrim("    Hello, World!")).toStrictEqual("Hello, World!");
    expect(ltrim("\n  \t Hello, World!")).toStrictEqual("Hello, World!");
    expect(ltrim("\0 \vHello, World!")).toStrictEqual("Hello, World!");
  });

  it("should not trim the right side of the string", () => {
    expect(ltrim("Hello, World!    ")).toStrictEqual("Hello, World!    ");
    expect(ltrim("    Hello, World!    ")).toStrictEqual("Hello, World!    ");
  });

  it("should return the string if there is nothing to be trimmed", () => {
    expect(ltrim("Hello, World!")).toStrictEqual("Hello, World!");
  });
});

describe("ltrim with different characters", () => {
  it("should return trimmed string with different characters", () => {
    expect(ltrim("rrrrrrrHello, World!", "r")).toStrictEqual("Hello, World!");
    expect(ltrim("rtrtrttrrHello, World!", "rt")).toStrictEqual("Hello, World!");
    expect(ltrim("0293174856Hello, World!", "0123456789")).toStrictEqual("Hello, World!");
  });

  it("should not trim the right side of the string", () => {
    expect(ltrim("Hello, World!rrrrrr", "r")).toStrictEqual("Hello, World!rrrrrr");
    expect(ltrim("rrrrrHello, World!rrrrr", "r")).toStrictEqual("Hello, World!rrrrr");
  });

  it("should trim regex special characters", () => {
    expect(ltrim(".*+Hello, World!", "*.+")).toStrictEqual("Hello, World!");
    expect(ltrim("[$$$*]Hello, World!", "*$[]")).toStrictEqual("Hello, World!");
  });
});

describe("ltrim with character range", () => {
  it("should return trimmed string for ranged characters", () => {
    expect(ltrim("\x09Hello, World!\x0A", "\x00..\x1F")).toStrictEqual("Hello, World!\x0A");
    expect(ltrim("0293174856Hello, World!", "0..9")).toStrictEqual("Hello, World!");
  });

  it("should also work with reverse range", () => {
    expect(ltrim("aiaxmioxjHello, World!", "z..a")).toStrictEqual("Hello, World!");    
  });

  it("should also work with multiple ranges", () => {
    expect(ltrim("38Hello, World!", "1..57..9")).toStrictEqual("Hello, World!");
    expect(ltrim("6Hello, World!", "1..57..9")).toStrictEqual("6Hello, World!");
  });
});