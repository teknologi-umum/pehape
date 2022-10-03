# frozen_string_literal: true

require "digest"

module Pehape
  def self.sha1(string, binary=nil)
    result = Digest::SHA1.new << string.to_s
    binary ? result.digest : result.hexdigest
  end
end
