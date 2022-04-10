# frozen_string_literal: false

module Pehape
  def self.ucfirst(str)
    dup = str
    dup[0] = dup[0].capitalize
    dup
  end
end
