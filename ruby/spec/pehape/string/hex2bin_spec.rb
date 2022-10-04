# frozen_string_literal: true

RSpec.describe Pehape do
  it "returns the binary representation of the given data or false on failure." do
    expect(Pehape.hex2bin("48656c6c6f20776f726c6421")).to eq("Hello world!")
  end
end
