# frozen_string_literal: true

require "digest"

module Pehape
  def self.md5_file(filename, binary=nil)
    return false unless File.exist?(filename.to_s)

    result = Digest::MD5.file(filename.to_s)
    binary ? result.digest : result.hexdigest
  end
end
