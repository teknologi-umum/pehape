# frozen_string_literal: true

RSpec.describe Pehape do
  it "returns the hexadecimal representation of the given string." do
    expect(Pehape.bin2hex("Hello world!")).to eq("48656c6c6f20776f726c6421")
  end
end
