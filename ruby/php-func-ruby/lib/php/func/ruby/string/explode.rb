# frozen_string_literal: true

module Php
  module Func
    module Ruby
        def self.explode(separator,string,limit = nil)
            if  separator.nil?
                raise StandardError.new "Separator cannot be nil"
            end 
            if  string.nil?
                raise StandardError.new "String cannot be nil"
            end
        # if  limit.class != Integer
        #     raise StandardError.new "Must Be Number" 
        # end
            case
            when limit.nil?
                return string.split(separator)
            when limit > 0
                return string.split(separator,limit)
            when limit < 0
                return string.split(separator,-limit)
            else
                return string.split(separator,limit) 
            end
        end
    end
  end
end