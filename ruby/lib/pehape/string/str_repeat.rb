# frozen_string_literal: false

module Pehape
  def self.str_repeat(string, times)
    begin
        string * times
    rescue => exception
        ""
    end
  end
end
