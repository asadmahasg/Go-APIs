package be

const (
	// Reserved
	RC_E_NO_BODY   RespCode = 2004
	RC_E_MALFORMED RespCode = 2005
	RC_E_RATELIMIT RespCode = 2010

	// POST /api/be/first
	RC_FIRST      RespCode = 1000
	RC_FIRST_FAIL RespCode = 2000
	
	// POST /api/be/second
	RC_SECOND      RespCode = 1000
	RC_SECOND_FAIL RespCode = 2000
)
