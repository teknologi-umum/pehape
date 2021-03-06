# frozen_string_literal: false

RSpec.describe Pehape do
  it "converts first character to lowercase" do
    expect(Pehape.lcfirst("PeHaPe")).to eq("peHaPe")
    expect(Pehape.lcfirst("PEHAPE")).to eq("pEHAPE")
  end
  it "removes leading whitespace and converts first character to lowercase" do
    expect(Pehape.lcfirst(" PeHaPe Hello")).to eq(" PeHaPe Hello")
    expect(Pehape.lcfirst("  PeHaPe  HE")).to eq("  PeHaPe  HE")
  end

  it "removes leading and trailing whitespace and converts first character to lowercase" do
    expect(Pehape.lcfirst(" PeHaPe Hello   ")).to eq(" PeHaPe Hello   ")
    expect(Pehape.lcfirst("  PeHaPe  HE   ")).to eq("  PeHaPe  HE   ")
  end

  it "converts non-english first character to lowercase" do
    expect(Pehape.lcfirst("_PeHaPe Hello")).to eq("_PeHaPe Hello")
    expect(Pehape.lcfirst("1PeHaPe  HE")).to eq("1PeHaPe  HE")
    expect(Pehape.lcfirst("ÃPeHaPe  HE")).to eq("ãPeHaPe  HE")
    expect(Pehape.lcfirst("ĐPeHaPe  HE")).to eq("đPeHaPe  HE")
  end
end
