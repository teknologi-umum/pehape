# frozen_string_literal: false

module Pehape
  def self.lcfirst(str)
    dup = str
    dup[0] = dup[0].downcase
    dup
  end
end
