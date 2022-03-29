# frozen_string_literal: true

RSpec.describe Pehape do
  it "implode array number" do
    number = [10, 20, 30, 40, 50, 60]
    expect(Pehape.implode("*", number)).to eq("10*20*30*40*50*60")
  end

  it "implode array string" do
    string = ["Pe", "Ha", "Pe"]
    expect(Pehape.implode(",", string)).to eq("Pe,Ha,Pe")
  end

  it "implode array random type data" do
    random = [true, false, 0, 1, "Pe"]
    expect(Pehape.implode("-", random)).to eq("true-false-0-1-Pe")
  end
  it "implode array nested array" do
    nested = [["Pe", "Ha"], "Pe"]
    expect(Pehape.implode("_", nested)).to eq("Pe_Ha_Pe")
  end
end
