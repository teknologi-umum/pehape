# frozen_string_literal: true

RSpec.describe Pehape do
  it "returns the repeated string" do
    expect(Pehape.str_repeat("hello world!", 3)).to eq("hello world!hello world!hello world!")
    expect(Pehape.str_repeat("hello world!", 0)).to eq("")
    expect(Pehape.str_repeat("hello world!", -3)).to eq("")
  end
end
