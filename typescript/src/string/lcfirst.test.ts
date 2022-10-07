import { lcfirst } from "./lcfirst";

const text = "Shizumu you ni tokete yuku you ni.";

describe("valid parameters", () => {
  it("should return lowercased first characters from capital text", () => {
    expect(lcfirst(text)).toEqual("shizumu you ni tokete yuku you ni.");
  });

  it("should return lowercased first characters from all uppercased text", () => {
    expect(lcfirst(text.toUpperCase())).toEqual("sHIZUMU YOU NI TOKETE YUKU YOU NI.");
  });

  it("should return lowercased first characters from all lowercased text", () => {
    expect(lcfirst(text.toLowerCase())).toEqual("shizumu you ni tokete yuku you ni.");
  });
});