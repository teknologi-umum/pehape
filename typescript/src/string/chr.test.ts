import { chr } from "./chr";

describe("chr", () => {
  it("ascii decimal value chr", () => {
    expect(chr(52)).toEqual("4");
  });

  it("ascii octal value chr", () => {
    expect(chr(0o52)).toEqual("*");
  });

  it("ascii hex value chr", () => {
    expect(chr(0x52)).toEqual("R");
  });
});
