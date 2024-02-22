# frozen_string_literal: true

RSpec.describe Pehape do
  describe "md5" do
    it "returns the hash as a 32-character hexadecimal number." do
      expect(Pehape.md5("Hello World!")).to eq("ed076287532e86365e841e92bfc50d8c")
      expect(Pehape.md5("apple")).to eq("1f3870be274f6c49b3e31a0c6728957f")
    end
  end
end
