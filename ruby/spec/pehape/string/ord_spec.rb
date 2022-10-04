# frozen_string_literal: false

RSpec.describe Pehape do
  describe "ord" do
    it "returns an integer between 0 and 255." do
      expect(Pehape.ord("Hello World!")).to eq(72)
      expect(Pehape.ord("\n")).to eq(10)
      expect(Pehape.ord("🐘")).to eq(128_024)
    end
  end
end
