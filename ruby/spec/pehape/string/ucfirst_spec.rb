# frozen_string_literal: true

RSpec.describe Pehape do
  it "first character uppercase" do
    expect(Pehape.ucfirst("peHaPe")).to eq("PeHaPe")
    expect(Pehape.ucfirst("PEHAPE")).to eq("PEHAPE")
  end
  it "remove leading whitespace and converts first character to uppercase" do
    expect(Pehape.ucfirst(" PeHaPe Hello")).to eq(" PeHaPe Hello")
    expect(Pehape.ucfirst("  peHaPe  HE")).to eq("  peHaPe  HE")
  end

  it "removes leading and trailing whitespace and converts first character to uppercase" do
    expect(Pehape.ucfirst(" PeHaPe Hello   ")).to eq(" PeHaPe Hello   ")
    expect(Pehape.ucfirst("  PeHaPe  HE   ")).to eq("  PeHaPe  HE   ")
  end

  it "converts non-english first character to  uppercase" do
    expect(Pehape.ucfirst("_PeHaPe Hello")).to eq("_PeHaPe Hello")
    expect(Pehape.ucfirst("1PeHaPe  HE")).to eq("1PeHaPe  HE")
    expect(Pehape.ucfirst("àPeHaPe  HE")).to eq("àPeHaPe  HE")
    expect(Pehape.ucfirst("ãPeHaPe  HE")).to eq("ãPeHaPe  HE")
    expect(Pehape.ucfirst("đPeHaPe  HE")).to eq("đPeHaPe  HE")
    expect(Pehape.ucfirst("×PeHaPe  HE")).to eq("×PeHaPe  HE")
  end
end
