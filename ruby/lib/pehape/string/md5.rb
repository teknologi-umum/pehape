# frozen_string_literal: true

require "digest"

module Pehape
  def self.md5(string, binary=nil)
    result = Digest::MD5.new << string.to_s
    binary ? result.digest : result.hexdigest
  end
end
