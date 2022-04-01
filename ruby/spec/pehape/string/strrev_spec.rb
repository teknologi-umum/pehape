# frozen_string_literal: true

RSpec.describe Pehape do
  it "reverse string space" do
    expect(Pehape.strrev("Hello World!")).to eq("!dlroW olleH")
  end

  it "reverse string double space" do
    expect(Pehape.strrev("Hello  World!")).to eq("!dlroW  olleH")
  end
end
