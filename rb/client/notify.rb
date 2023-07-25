$LOAD_PATH.unshift(File.expand_path('../../', __FILE__))
require 'grpc'
require 'proto/foobar_services_pb'

def main
  stub = Proto::FoobarService::Stub.new('win11:8080', :this_channel_is_insecure)

  # Create a request
  req = Proto::RescanRequest.new

  # Send the request to the server
  res = stub.trigger_rescan(req)

  # Print the response
  puts "Success: #{res.success}"
end

main
