# frozen_string_literal: false

module Pehape
  def self.str_rot13(string)
    result = string.to_s.chars.map do |char|
      new_char = case char
                 when "a".."m", "A".."M"
                   char.ord + 13
                 when "n".."z", "N".."Z"
                   char.ord - 13
                 else
                   char.ord
                 end
      new_char.chr
    end
    result.join
  end
end
