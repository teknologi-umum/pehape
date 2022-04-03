# frozen_string_literal: true

RSpec.describe Pehape do
  it "first character uppercase" do
    expect(Pehape.ucfirst("peHaPe")).to eq("PeHaPe")
    expect(Pehape.ucfirst("PEHAPE")).to eq("PEHAPE")
  end
  it "remove left whitespace and get first character uppercase" do
    expect(Pehape.ucfirst(" PeHaPe Hello")).to eq("PeHaPe Hello")
    expect(Pehape.ucfirst("  peHaPe  HE")).to eq("PeHaPe  HE")
  end

  it "remove left and right whitespace and get first character uppercase" do
    expect(Pehape.ucfirst(" PeHaPe Hello   ")).to eq("PeHaPe Hello")
    expect(Pehape.ucfirst("  PeHaPe  HE   ")).to eq("PeHaPe  HE")
  end

  it "another character except abjad and get first character uppercase" do
    expect(Pehape.ucfirst("_PeHaPe Hello")).to eq("_PeHaPe Hello")
    expect(Pehape.ucfirst("1PeHaPe  HE")).to eq("1PeHaPe  HE")
    expect(Pehape.ucfirst("àPeHaPe  HE")).to eq("ÀPeHaPe  HE")
    expect(Pehape.ucfirst("ãPeHaPe  HE")).to eq("ÃPeHaPe  HE")
    expect(Pehape.ucfirst("ÀPeHaPe  HE")).to eq("ÀPeHaPe  HE")
    expect(Pehape.ucfirst("đPeHaPe  HE")).to eq("ĐPeHaPe  HE")
    expect(Pehape.ucfirst("×PeHaPe  HE")).to eq("×PeHaPe  HE")
  end
end
