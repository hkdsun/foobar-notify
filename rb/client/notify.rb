$LOAD_PATH.unshift(File.expand_path("../../", __FILE__))
require "grpc"
require "proto/foobar_services_pb"

module FoobarNotify
  def self.notify
    stub = Proto::FoobarService::Stub.new("win11:9031", :this_channel_is_insecure)

    # Create a request
    req = Proto::RescanRequest.new

    # Send the request to the server
    stub.trigger_rescan(req)
  end
end
