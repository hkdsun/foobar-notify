$LOAD_PATH.unshift(File.expand_path('../../', __FILE__))
require 'grpc'
require 'proto/foobar_services_pb'

def main
  # Create a gRPC client
  stub = Proto::FoobarService::Stub.new('localhost:8080', :this_channel_is_insecure)

  # Create a request
  req = Proto::RescanRequest.new(foobar_path: "/path/to/foobar")

  # Send the request to the server
  res = stub.trigger_rescan(req)

  # Print the response
  puts "Success: #{res.success}"
end

main
