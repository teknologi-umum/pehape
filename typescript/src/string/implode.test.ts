import { implode } from "./implode";

const arrayOfString = ["The", "lazy", "brown", "fox"];

describe("empty parameters", () => {
  it("should treat separator as empty string if it is null", () => {
    expect(implode(null, arrayOfString)).toBe("Thelazybrownfox");
  });
  it("should return empty array if arrayOfString is empty", () => {
    expect(implode(null, [])).toBe("");
  });
});

describe("separator is white space", () => {
  it("should join array of strings with whitespace", () => {
    expect(implode(" ", arrayOfString)).toBe("The lazy brown fox");
  });
  it("should join array of strings with double whitespace", () => {
    expect(implode("  ", arrayOfString)).toBe("The  lazy  brown  fox");
  });
});

describe("separator is non-white space", () => {
  it("should join array of strings with words", () => {
    expect(implode("js", arrayOfString)).toBe("Thejslazyjsbrownjsfox");
  });
  it("should join array of strings with special characters", () => {
    expect(implode("\t", arrayOfString)).toBe("The\tlazy\tbrown\tfox");
  });
});
