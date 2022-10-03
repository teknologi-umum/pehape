# frozen_string_literal: true

RSpec.describe Pehape do
  describe "md5_file" do
    it "returns a string on success." do
      filename = File.expand_path("./README.md")
      expect(Pehape.md5_file(filename)).to eq("6cccfa61af5a5310f08f3fabcbe230ae")
    end

    it "returns false on fail." do
      filename = File.expand_path("./robot.txt")
      expect(Pehape.md5_file(filename)).to eq(false)
    end
  end
end
