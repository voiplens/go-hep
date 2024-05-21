package binary

const ChunkHeaderLength = 6

const (
	VendorGeneric    uint16 = 0x0000 // No specific vendor, generic chunk types, see above
	VendorFreeSWITCH uint16 = 0x0001 // FreeSWITCH (www.freeswitch.org)
	VendorKamailio   uint16 = 0x0002 // Kamailio/SER (www.kamailio.org)
	VendorOpenSIPS   uint16 = 0x0003 // OpenSIPS (www.opensips.org)
	VendorAsterisk   uint16 = 0x0004 // Asterisk (www.asterisk.org)
	VendorHomer      uint16 = 0x0005 // Homer Project (www.sipcapture.org)
	VendorSipXecs    uint16 = 0x0006 // SipXecs (www.sipfoundry.org)
	VendorYetiSwitch uint16 = 0x0007 // Yeti Switch (yeti-switch.org)
	VendoreGenesys   uint16 = 0x0008 // Genesys (www.genesys.com)
)

// HEP chuncks
const (
	TypeIPProtocolFamily  uint16 = 0x0001 // uint8 IP protocol family
	TypeIPProtocolID      uint16 = 0x0002 // uint8 IP protocol ID
	TypeIP4SrcIP          uint16 = 0x0003 // inet4-addr IPv4 source address
	TypeIP4DstIP          uint16 = 0x0004 // inet4-addr IPv4 destination address
	TypeIP6SrcIP          uint16 = 0x0005 // inet6-addr IPv6 source address
	TypeIP6DstIP          uint16 = 0x0006 // inet6-addr IPv6 destination address
	TypeSrcPort           uint16 = 0x0007 // uint16 protocol source port (UDP, TCP, SCTP)
	TypeDstPort           uint16 = 0x0008 // uint16 protocol destination port (UDP, TCP, SCTP)
	TypeTsec              uint16 = 0x0009 // uint32 timestamp, seconds since 01/01/1970 (epoch)
	TypeTmsec             uint16 = 0x000a // uint32 timestamp microseconds offset (added to timestamp)
	TypeProtoType         uint16 = 0x000b // uint8 protocol type (SIP/H323/RTP/MGCP/M2UA)
	TypeNodeID            uint16 = 0x000c // uint32 capture agent ID (202, 1201, 2033...)
	TypeAliveTimSec       uint16 = 0x000d // uint16 keep alive timer (sec)
	TypeNodePW            uint16 = 0x000e // octet-string authenticate key (plain text / TLS connection)
	TypePayload           uint16 = 0x000f // octet-string captured packet payload
	TypeCompressedPayload uint16 = 0x0010 // octet-string captured compressed payload (gzip/inflate)
	TypeCID               uint16 = 0x0011 // octet-string Internal correlation id
	TypeVlanID            uint16 = 0x0012 // uint16 Vlan ID
	TypeNodeName          uint16 = 0x0013 // octet-string capture agent ID (“node1”, “node2”, “node3”...)
	TypeSrcMAC            uint16 = 0x0014 // uint64 Source MAC
	TypeDstMAC            uint16 = 0x0015 // uint64 Destination MAC
	TypeEthernetType      uint16 = 0x0016 // uint16 Ethernet Type
	TypeTCPFlag           uint16 = 0x0017 // uint8 TCP Flag [SYN.PUSH...]
	TypeIPTos             uint16 = 0x0018 // uint8 IP TOS
	TypeMOS               uint16 = 0x0020 // uint16 MOS value
	TypeRFactor           uint16 = 0x0021 // uint16 R-Factor
	TypeGeoLocation       uint16 = 0x0022 // octet-string GEO Location
	TypeJitter            uint16 = 0x0023 // uint32 Jitter
	TypeTransactionType   uint16 = 0x0024 // octet-string Transaction type [call, registration]
	TypePayloadJSONKeys   uint16 = 0x0025 // octet-string Payload JSON Keys
	TypeTagsValue         uint16 = 0x0026 // octet-string Tags’ values
	TypeTagType           uint16 = 0x0027 // uint16 Type of tag
	TypeEventType         uint16 = 0x0028 // uint16 Event type [recording|interception|
	TypeGroupID           uint16 = 0x0029 // octet-string Group ID
)
