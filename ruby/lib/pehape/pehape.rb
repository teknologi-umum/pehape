# frozen_string_literal: true

require_relative "pehape/version"
require_relative "pehape/string/index"

begin
  require "pry"
rescue LoadError
end

module Pehape
  class Error < StandardError; end

  def self.aksi
    "Aksi"
  end
end
