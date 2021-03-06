# gRPC サービス仕様書

## インデックス

- サービス仕様書
  - [Services](#Services)
  
    
    - [protobuf/ping.proto](#protobuf/ping.proto)
    
      - [Ping](#ping.Ping)
    
  

  - [Messages](#Messages)
  
    
    - [protobuf/ping.proto](#protobuf/ping.proto)
    
      - [PingResponse](#ping.PingResponse)
    
  

  - [Enums](#Enums)
  

  - [Extensions](#Extensions)
  

- [スカラー値型](#スカラー値型)

## API仕様

### Services

#### protobuf/ping.proto


- Ping

  | Method Name | Request Type | Response Type | Description |
  | ----------- | ------------ | ------------- | ------------|
  | Ping | [.google.protobuf.Empty](#google.protobuf.Empty) | [PingResponse](#ping.PingResponse) |  |
  

 <!-- end services -->

### Messages

#### protobuf/ping.proto


- PingResponse

  
  | Field | Type | Label | Description |
  | ----- | ---- | ----- | ----------- |
  | message | [string](#string) |  |  |
  



<!-- end messages -->

### Enums
<!-- end enums -->

### File-level Extensions
 <!-- end HasExtensions -->

## スカラー値型

| .proto Type | Notes | Go Type | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | -------- | --------- | ----------- |
| <a name="double" /> double |  | float64 | double | double | float |
| <a name="float" /> float |  | float32 | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | []byte | string | ByteString | str |
