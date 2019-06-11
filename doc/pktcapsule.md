# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [pktcapsule.proto](#pktcapsule.proto)
    - [DecapsulateRequest](#pktcapsule.DecapsulateRequest)
    - [EncapsulateRequest](#pktcapsule.EncapsulateRequest)
    - [Packet](#pktcapsule.Packet)
  
  
  
    - [PktCapsule](#pktcapsule.PktCapsule)
  

- [Scalar Value Types](#scalar-value-types)



<a name="pktcapsule.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pktcapsule.proto
pktcapsule.proto

pktcapsule is the server to encapsulate or decapsulate packet data with
IPinIP.
Currently, we support IPv4 only.


<a name="pktcapsule.DecapsulateRequest"></a>

### DecapsulateRequest
Message DecapsulateRequest represents the request message for Decapsulate
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  | Packet data which will be decapsulated. Linklayer header must be removed. |






<a name="pktcapsule.EncapsulateRequest"></a>

### EncapsulateRequest
Message EncapsulateRequest represents the request message for Encapsulate
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  | Packet data which will be encapsulated. Linklayer header must be removed. |
| src_ip | [fixed32](#fixed32) |  | Source IP address which will be used in newly added IP header. |
| dst_ip | [fixed32](#fixed32) |  | Destination IP address which will be used in newly added IP header. |






<a name="pktcapsule.Packet"></a>

### Packet
Message Packet represents the packet data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  | Encapsulated/Decapsulated packet data. |





 

 

 


<a name="pktcapsule.PktCapsule"></a>

### PktCapsule
Service PktCapsule encapsulates or decapsulates given packet data.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Encapsulate | [EncapsulateRequest](#pktcapsule.EncapsulateRequest) | [Packet](#pktcapsule.Packet) | Encapsulate adds newly IP header to the given packet data. |
| Decapsulate | [DecapsulateRequest](#pktcapsule.DecapsulateRequest) | [Packet](#pktcapsule.Packet) | Decapsulate removes outer IP header from the given packet data. |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

