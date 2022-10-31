import { rtrim } from "./rtrim";

describe("rtrim with default characters", () => {
  it("should return trimmed string", () => {
    expect(rtrim("Hello, World!    ")).toStrictEqual("Hello, World!");
    expect(rtrim("Hello, World!\r\n  \t ")).toStrictEqual("Hello, World!");
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
    expect(rtrim("Hello, World!.*+", "*.+")).toStrictEqual("Hello, World!");
    expect(rtrim("Hello, World![$$$*]", "*$[]")).toStrictEqual("Hello, World!");
  });
});