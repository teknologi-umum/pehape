# frozen_string_literal: true

RSpec.describe Pehape do
  it "explode space string separate" do
    str = "Pe ha pe"
    expect(Pehape.explode(" ", str)).to include("Pe", "ha", "pe")
    expect(Pehape.explode(" ", str, 0)).to include("Pe ha pe")
    expect(Pehape.explode(" ", str, -1)).to include("Pe ha pe")
    expect(Pehape.explode(" ", str, -2)).to include("Pe", "ha pe")
    expect(Pehape.explode(" ", str, -3)).to include("Pe", "ha", "pe")
    expect(Pehape.explode(" ", str, 1)).to include("Pe ha pe")
    expect(Pehape.explode(" ", str, 2)).to include("Pe", "ha pe")
    expect(Pehape.explode(" ", str, 3)).to include("Pe", "ha", "pe")
  end

  it "explode space string separate by double space" do
    str = "Pe ha pe"
    expect(Pehape.explode("  ", str)).to include("Pe ha pe")
  end

  it "explode double space string separate by space" do
    str = "Pe  ha  pe"
    expect(Pehape.explode(" ", str)).to include("Pe", "ha", "pe")
  end

  it "explode nil separator raise error argument" do
    str = "Pe  ha  pe"
    expect { Pehape.explode(nil, str) }.to raise_error(ArgumentError, "Separator cannot be nil")
  end
end
