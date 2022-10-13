import { lcfirst } from "./lcfirst";

const text = "Shizumu you ni tokete yuku you ni.";

describe("parameter is empty string", () => {
  it("should return empty string", () => {
    expect(lcfirst("")).toBe("");
  });
});

describe("parameter is whitespace", () => {
  it("should return the whitespace", () => {
    expect(lcfirst("  ")).toBe("  ");
  });
});

describe("parameter not empty string", () => {
  it("should return lowercased first characters from capital text", () => {
    expect(lcfirst(text)).toEqual("shizumu you ni tokete yuku you ni.");
  });

  it("should return lowercased first characters from all uppercased text", () => {
    expect(lcfirst(text.toUpperCase())).toEqual("sHIZUMU YOU NI TOKETE YUKU YOU NI.");
  });

  it("should return lowercased first characters from all lowercased text", () => {
    expect(lcfirst(text.toLowerCase())).toEqual("shizumu you ni tokete yuku you ni.");
  });

  it("should return the same for non alphabetical characters", () => {
    expect(lcfirst("沈むように溶けてゆくように")).toEqual("沈むように溶けてゆくように");
  });

  it("should return lowercased first characters from alphabetical with special characters", () => {
    expect(lcfirst("Éxample")).toEqual("éxample");
  });

  it("should return the same as input", () => {
    expect(lcfirst("123456789")).toBe("123456789");
    expect(lcfirst("!@#$%%")).toBe("!@#$%%");
  });
});