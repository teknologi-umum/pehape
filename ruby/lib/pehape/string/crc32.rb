# frozen_string_literal: true

require "zlib"

module Pehape
  def self.crc32(string)
    Zlib.crc32(string.to_s)
  end
end
