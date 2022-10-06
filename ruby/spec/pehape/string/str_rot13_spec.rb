# frozen_string_literal: false

RSpec.describe Pehape do
  describe "str_rot13" do
    it "returns the ROT13 version of the given string." do
      expect(Pehape.str_rot13("PHP 4.3.0")).to eq("CUC 4.3.0")
      expect(Pehape.str_rot13("Hello World!")).to eq("Uryyb Jbeyq!")
    end
  end
end
