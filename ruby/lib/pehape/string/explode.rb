# frozen_string_literal: true

module Pehape
  def self.explode(separator, string, limit=nil)
    raise ArgumentError.new("Separator cannot be nil") if separator.nil?
    raise ArgumentError.new("String cannot be nil") if string.nil?

    if limit.nil?
      string.split(separator)
    elsif limit.positive?
      string.split(separator, limit)
    elsif limit.negative?
      string.split(separator, -limit)
    else
      string.split(separator, 1)
    end
  end
end
