# frozen_string_literal: true

module Pehape
  def self.bin2hex(string=nil)
    string.to_s.unpack1("H*")
  end
end
