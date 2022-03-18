# frozen_string_literal: true
require_relative "ruby/version"
require_relative "ruby/string/index"

begin
  require "pry"
rescue LoadError
end

module Php
  module Func
    module Ruby
      class Error < StandardError; end

      def self.aksi
        "Aksi"
      end
      
    end
  end
end
