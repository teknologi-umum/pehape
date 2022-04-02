# frozen_string_literal: true

module Pehape
  def self.ucfirst(str)
    character = str.strip
    character[0] = character[0].capitalize
    character
  end
end
