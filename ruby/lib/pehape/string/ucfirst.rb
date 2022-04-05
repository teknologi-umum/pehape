# frozen_string_literal: true

module Pehape
  def self.ucfirst(str)
    str.sub(/^(\w)/, &:capitalize)
  end
end
