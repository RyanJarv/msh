#!/usr/local/bin/msh lambda ruby

require 'json'

def lambda_handler(event:, context:)
  puts "Hello #{ENV["MSH_WHERE"]} World!"

  { statusCode: 200 "message": "OK"}
end
