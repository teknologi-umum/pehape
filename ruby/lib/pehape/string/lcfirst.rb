# frozen_string_literal: true

module Pehape
  def self.lcfirst(str)
    character = str.strip
    character[0] = character[0].downcase
    character
  end
end
