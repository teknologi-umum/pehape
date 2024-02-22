# frozen_string_literal: true

RSpec.describe Pehape do
  describe "crc32" do
    it "returns the crc32 checksum of string as an integer." do
      expect(Pehape.crc32("Hello World!")).to eq(472_456_355)
    end
  end
end
