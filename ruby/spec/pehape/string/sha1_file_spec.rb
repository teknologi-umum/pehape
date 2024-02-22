# frozen_string_literal: true

RSpec.describe Pehape do
  describe "sha1_file" do
    it "returns file hash as a string on success." do
      filename = File.expand_path("./README.md")
      expect(Pehape.sha1_file(filename)).to eq("64285a71bb456bd05a5466ff7076bf2c377ecd57")
    end

    it "returns false on fail." do
      filename = File.expand_path("./robot.txt")
      expect(Pehape.sha1_file(filename)).to eq(false)
    end
  end
end
