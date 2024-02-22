# frozen_string_literal: true

RSpec.describe Pehape do
  describe "sha1" do
    it "returns the sha1 hash as a string." do
      expect(Pehape.sha1("Hello World!")).to eq("2ef7bde608ce5404e97d5f042f95f89f1c232871")
    end
  end
end
