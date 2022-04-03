# frozen_string_literal: true

RSpec.describe Pehape do
  it "first character lowercase" do
    expect(Pehape.lcfirst("PeHaPe")).to eq("peHaPe")
    expect(Pehape.lcfirst("PEHAPE")).to eq("pEHAPE")
  end
  it "remove left whitespace and get first character lowercase" do
    expect(Pehape.lcfirst(" PeHaPe Hello")).to eq("peHaPe Hello")
    expect(Pehape.lcfirst("  PeHaPe  HE")).to eq("peHaPe  HE")
  end

  it "remove left and right whitespace and get first character lowercase" do
    expect(Pehape.lcfirst(" PeHaPe Hello   ")).to eq("peHaPe Hello")
    expect(Pehape.lcfirst("  PeHaPe  HE   ")).to eq("peHaPe  HE")
  end

  it "another character except abjad and get first character lowercase" do
    expect(Pehape.lcfirst("_PeHaPe Hello")).to eq("_PeHaPe Hello")
    expect(Pehape.lcfirst("1PeHaPe  HE")).to eq("1PeHaPe  HE")
    expect(Pehape.lcfirst("àPeHaPe  HE")).to eq("àPeHaPe  HE")
    expect(Pehape.lcfirst("ãPeHaPe  HE")).to eq("ãPeHaPe  HE")
    expect(Pehape.lcfirst("ÀPeHaPe  HE")).to eq("àPeHaPe  HE")
    expect(Pehape.lcfirst("ĐPeHaPe  HE")).to eq("đPeHaPe  HE")
    expect(Pehape.lcfirst("×PeHaPe  HE")).to eq("×PeHaPe  HE")
  end
end
