# frozen_string_literal: true
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/foobar.proto

require 'google/protobuf'


descriptor_data = "\n\x12proto/foobar.proto\x12\x06\x66oobar\"$\n\rRescanRequest\x12\x13\n\x0b\x66oobar_path\x18\x01 \x01(\t\"!\n\x0eRescanResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x32Q\n\rFoobarService\x12@\n\rTriggerRescan\x12\x15.foobar.RescanRequest\x1a\x16.foobar.RescanResponse\"\x00\x42\x11Z\x07./proto\xea\x02\x05Protob\x06proto3"

pool = Google::Protobuf::DescriptorPool.generated_pool

begin
  pool.add_serialized_file(descriptor_data)
rescue TypeError => e
  # Compatibility code: will be removed in the next major version.
  require 'google/protobuf/descriptor_pb'
  parsed = Google::Protobuf::FileDescriptorProto.decode(descriptor_data)
  parsed.clear_dependency
  serialized = parsed.class.encode(parsed)
  file = pool.add_serialized_file(serialized)
  warn "Warning: Protobuf detected an import path issue while loading generated file #{__FILE__}"
  imports = [
  ]
  imports.each do |type_name, expected_filename|
    import_file = pool.lookup(type_name).file_descriptor
    if import_file.name != expected_filename
      warn "- #{file.name} imports #{expected_filename}, but that import was loaded as #{import_file.name}"
    end
  end
  warn "Each proto file must use a consistent fully-qualified name."
  warn "This will become an error in the next major version."
end

module Proto
  RescanRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("foobar.RescanRequest").msgclass
  RescanResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("foobar.RescanResponse").msgclass
end
