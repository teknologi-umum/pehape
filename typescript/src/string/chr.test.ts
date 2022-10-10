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
  it("ascii the decimal number -1 value chr", () => {
    expect(chr(-1)).toEqual("￿");
  });
  it("ascii the decimal number 255 value chr", () => {
    expect(chr(255)).toEqual("ÿ");
  });
});
