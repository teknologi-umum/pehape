# frozen_string_literal: false

module Pehape
  def self.str_repeat(string, times)
    times = times.to_i
    return "" if times <= 0

    string * times
  end
end
