# frozen_string_literal: true

RSpec.describe Pehape do
  it "first character lowercase" do
    expect(Pehape.lcfirst("PeHaPe")).to eq("peHaPe")
    expect(Pehape.lcfirst("PEHAPE")).to eq("pEHAPE")
  end
  it "remove leading whitespace and get first character to lowercase" do
    expect(Pehape.lcfirst(" PeHaPe Hello")).to eq(" PeHaPe Hello")
    expect(Pehape.lcfirst("  PeHaPe  HE")).to eq("  PeHaPe  HE")
  end

  it "remove leading and trailing whitespace and converts first character to lowercase" do
    expect(Pehape.lcfirst(" PeHaPe Hello   ")).to eq(" PeHaPe Hello   ")
    expect(Pehape.lcfirst("  PeHaPe  HE   ")).to eq("  PeHaPe  HE   ")
  end

  it "converts non-english first character to lowercase" do
    expect(Pehape.lcfirst("_PeHaPe Hello")).to eq("_PeHaPe Hello")
    expect(Pehape.lcfirst("1PeHaPe  HE")).to eq("1PeHaPe  HE")
    expect(Pehape.lcfirst("àPeHaPe  HE")).to eq("àPeHaPe  HE")
    expect(Pehape.lcfirst("ãPeHaPe  HE")).to eq("ãPeHaPe  HE")
    expect(Pehape.lcfirst("đPeHaPe  HE")).to eq("đPeHaPe  HE")
    expect(Pehape.lcfirst("×PeHaPe  HE")).to eq("×PeHaPe  HE")
  end
end
