# frozen_string_literal: true

RSpec.describe Pehape do
  it "ascii decimal value chr" do
    expect(Pehape.chr(52)).to eq("4")
  end

  it "ascii octal value chr" do
    expect(Pehape.chr(0o52)).to eq("*")
  end

  it "ascii hex value chr" do
    expect(Pehape.chr(0x52)).to eq("R")
  end
end
