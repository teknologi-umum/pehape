# frozen_string_literal: true

module Pehape
  def self.lcfirst(str)
    str.sub(/^(\w)/, &:downcase)
  end
end
