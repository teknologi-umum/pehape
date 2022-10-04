# frozen_string_literal: true

module Pehape
  def self.hex2bin(string=nil)
    [string.to_s].pack("H*")
  end
end
