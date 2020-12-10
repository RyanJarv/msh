#!/usr/bin/env ruby

require 'json'

def lambda_handler(event:, context:)
  # Only run once if we're running in docker
  if File.exists? "/.dockerenv"
    Thread.new { sleep(0.3); Process.kill("INT", Process.ppid) }
  end

  { statusCode: 200, message: "hello from lambda compose" }
end