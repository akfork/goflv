package flv

var (
	HEADER_BYTES = []byte{'F', 'L', 'V', 0x01, 0x05, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, // 13 bytes, flv header
		0x12, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 11 bytes, scritp tag
		0x02, 0x00, 0x0a, 'o', 'n', 'M', 'e', 't', 'a', 'D', 'a', 't', 'a', // 13 bytes, metadata content, 1st amf
		0x08, 0x00, 0x00, 0x00, 0x01, // 5 bytes, 2dn amf header
		0x00, 0x08, 'd', 'u', 'r', 'a', 't', 'i', 'o', 'n', // 10 bytes, duration
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 9 bytes, duration value
		0x00, 0x00, 0x09, // 3 bytes, AMF_END_OF_OBJECT
		0x00, 0x00, 0x00, 0x33} // 4 bytes, previous tag size
)

const (
	AUDIO_TAG       = byte(0x08)
	VIDEO_TAG       = byte(0x09)
	SCRIPT_DATA_TAG = byte(0x12)
	DURATION_OFFSET = 53
	HEADER_LEN      = 13
)
