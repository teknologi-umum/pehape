# frozen_string_literal: false

RSpec.describe Pehape do
  it "converts first character to uppercase" do
    expect(Pehape.ucfirst("peHaPe")).to eq("PeHaPe")
    expect(Pehape.ucfirst("PEHAPE")).to eq("PEHAPE")
  end
  it "removes leading whitespace and converts first character to uppercase" do
    expect(Pehape.ucfirst(" PeHaPe Hello")).to eq(" PeHaPe Hello")
    expect(Pehape.ucfirst("  peHaPe  HE")).to eq("  peHaPe  HE")
  end

  it "removes leading and trailing whitespace and converts first character to uppercase" do
    expect(Pehape.ucfirst(" PeHaPe Hello   ")).to eq(" PeHaPe Hello   ")
    expect(Pehape.ucfirst("  PeHaPe  HE   ")).to eq("  PeHaPe  HE   ")
  end

  it "converts non-english first character to uppercase" do
    expect(Pehape.ucfirst("_PeHaPe Hello")).to eq("_PeHaPe Hello")
    expect(Pehape.ucfirst("1PeHaPe  HE")).to eq("1PeHaPe  HE")
    expect(Pehape.ucfirst("ãPeHaPe  HE")).to eq("ÃPeHaPe  HE")
    expect(Pehape.ucfirst("đPeHaPe  HE")).to eq("ĐPeHaPe  HE")
  end
end
