package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// ApplicationLog application log
// swagger:model ApplicationLog
type ApplicationLog struct {

	// Placeholder for description of property adf of obj type ApplicationLog field type str  type boolean
	// Required: true
	Adf bool `json:"adf"`

	// all_request_headers of ApplicationLog.
	AllRequestHeaders string `json:"all_request_headers,omitempty"`

	// all_response_headers of ApplicationLog.
	AllResponseHeaders string `json:"all_response_headers,omitempty"`

	// Number of app_response_time.
	AppResponseTime int64 `json:"app_response_time,omitempty"`

	//  Enum options - NOT_UPDATED, BY_CONTENT_REWRITE_PROFILE, BY_DATA_SCRIPT. Field introduced in 17.1.1.
	BodyUpdated string `json:"body_updated,omitempty"`

	// Placeholder for description of property cache_hit of obj type ApplicationLog field type str  type boolean
	CacheHit bool `json:"cache_hit,omitempty"`

	// Placeholder for description of property cacheable of obj type ApplicationLog field type str  type boolean
	Cacheable bool `json:"cacheable,omitempty"`

	// client_browser of ApplicationLog.
	ClientBrowser string `json:"client_browser,omitempty"`

	// Number of client_dest_port.
	// Required: true
	ClientDestPort int32 `json:"client_dest_port"`

	// client_device of ApplicationLog.
	ClientDevice string `json:"client_device,omitempty"`

	//  Enum options - INSIGHTS_DISABLED, NO_INSIGHTS_NOT_SAMPLED_COUNT, NO_INSIGHTS_NOT_SAMPLED_TYPE, NO_INSIGHTS_NOT_SAMPLED_SKIP_URI, NO_INSIGHTS_NOT_SAMPLED_URI_NOT_IN_LIST, NO_INSIGHTS_NOT_SAMPLED_CLIENT_IP_NOT_IN_RANGE, NO_INSIGHTS_NOT_SAMPLED_OTHER, ACTIVE_INSIGHTS_FAILED, ACTIVE_INSIGHTS_ENABLED, PASSIVE_INSIGHTS_ENABLED.
	ClientInsights string `json:"client_insights,omitempty"`

	// Number of client_ip.
	// Required: true
	ClientIP int32 `json:"client_ip"`

	// client_location of ApplicationLog.
	ClientLocation string `json:"client_location,omitempty"`

	// client_os of ApplicationLog.
	ClientOs string `json:"client_os,omitempty"`

	// Number of client_rtt.
	// Required: true
	ClientRtt int32 `json:"client_rtt"`

	// Number of client_src_port.
	// Required: true
	ClientSrcPort int32 `json:"client_src_port"`

	//  Enum options - NO_COMPRESSION_DISABLED, NO_COMPRESSION_GZIP_CONTENT, NO_COMPRESSION_CONTENT_TYPE, NO_COMPRESSION_CUSTOM_FILTER, NO_COMPRESSION_AUTO_FILTER, NO_COMPRESSION_MIN_LENGTH, NO_COMPRESSION_CAN_BE_COMPRESSED, COMPRESSED.
	Compression string `json:"compression,omitempty"`

	// Number of compression_percentage.
	CompressionPercentage int32 `json:"compression_percentage,omitempty"`

	// Placeholder for description of property connection_error_info of obj type ApplicationLog field type str  type object
	ConnectionErrorInfo *ConnErrorInfo `json:"connection_error_info,omitempty"`

	// Number of data_transfer_time.
	DataTransferTime int64 `json:"data_transfer_time,omitempty"`

	// Placeholder for description of property datascript_error_trace of obj type ApplicationLog field type str  type object
	DatascriptErrorTrace *DataScriptErrorTrace `json:"datascript_error_trace,omitempty"`

	// Log created by the invocations of the DataScript api avi.vs.log().
	DatascriptLog string `json:"datascript_log,omitempty"`

	// etag of ApplicationLog.
	Etag string `json:"etag,omitempty"`

	// Response headers received from backend server.
	HeadersReceivedFromServer string `json:"headers_received_from_server,omitempty"`

	// Request headers sent to backend server.
	HeadersSentToServer string `json:"headers_sent_to_server,omitempty"`

	// host of ApplicationLog.
	Host string `json:"host,omitempty"`

	// http_request_policy_rule_name of ApplicationLog.
	HTTPRequestPolicyRuleName string `json:"http_request_policy_rule_name,omitempty"`

	// http_response_policy_rule_name of ApplicationLog.
	HTTPResponsePolicyRuleName string `json:"http_response_policy_rule_name,omitempty"`

	// http_security_policy_rule_name of ApplicationLog.
	HTTPSecurityPolicyRuleName string `json:"http_security_policy_rule_name,omitempty"`

	// http_version of ApplicationLog.
	HTTPVersion string `json:"http_version,omitempty"`

	// Number of log_id.
	// Required: true
	LogID int32 `json:"log_id"`

	// method of ApplicationLog.
	Method string `json:"method,omitempty"`

	// microservice of ApplicationLog.
	Microservice string `json:"microservice,omitempty"`

	// microservice_name of ApplicationLog.
	MicroserviceName string `json:"microservice_name,omitempty"`

	// network_security_policy_rule_name of ApplicationLog.
	NetworkSecurityPolicyRuleName string `json:"network_security_policy_rule_name,omitempty"`

	// Placeholder for description of property persistence_used of obj type ApplicationLog field type str  type boolean
	PersistenceUsed bool `json:"persistence_used,omitempty"`

	// Number of persistent_session_id.
	PersistentSessionID int64 `json:"persistent_session_id,omitempty"`

	// pool of ApplicationLog.
	Pool string `json:"pool,omitempty"`

	// pool_name of ApplicationLog.
	PoolName string `json:"pool_name,omitempty"`

	// redirected_uri of ApplicationLog.
	RedirectedURI string `json:"redirected_uri,omitempty"`

	// referer of ApplicationLog.
	Referer string `json:"referer,omitempty"`

	// Number of report_timestamp.
	// Required: true
	ReportTimestamp int64 `json:"report_timestamp"`

	// request_content_type of ApplicationLog.
	RequestContentType string `json:"request_content_type,omitempty"`

	// Number of request_headers.
	RequestHeaders int32 `json:"request_headers,omitempty"`

	// Number of request_length.
	RequestLength int64 `json:"request_length,omitempty"`

	//  Enum options - AVI_HTTP_REQUEST_STATE_CONN_ACCEPT, AVI_HTTP_REQUEST_STATE_WAITING_FOR_REQUEST, AVI_HTTP_REQUEST_STATE_SSL_HANDSHAKING, AVI_HTTP_REQUEST_STATE_PROCESSING_SPDY, AVI_HTTP_REQUEST_STATE_READ_CLIENT_REQ_LINE, AVI_HTTP_REQUEST_STATE_READ_CLIENT_REQ_HDR, AVI_HTTP_REQUEST_STATE_CONNECT_TO_UPSTREAM, AVI_HTTP_REQUEST_STATE_SEND_REQ_TO_UPSTREAM, AVI_HTTP_REQUEST_STATE_READ_RESP_HDR_FROM_UPSTREAM, AVI_HTTP_REQUEST_STATE_SEND_TO_CLIENT, AVI_HTTP_REQUEST_STATE_KEEPALIVE, AVI_HTTP_REQUEST_STATE_PROXY_UPGRADED_CONN, AVI_HTTP_REQUEST_STATE_CLOSING_REQUEST, AVI_HTTP_REQUEST_STATE_READ_FROM_UPSTREAM, AVI_HTTP_REQUEST_STATE_READ_PROXY_PROTOCOL, AVI_HTTP_REQUEST_STATE_READ_CLIENT_PIPELINE_REQ_LINE, AVI_HTTP_REQUEST_STATE_SSL_HANDSHAKE_TO_UPSTREAM, AVI_HTTP_REQUEST_STATE_WAITING_IN_CONNPOOL_CACHE.
	RequestState string `json:"request_state,omitempty"`

	// Number of response_code.
	ResponseCode int32 `json:"response_code,omitempty"`

	// response_content_type of ApplicationLog.
	ResponseContentType string `json:"response_content_type,omitempty"`

	// Number of response_headers.
	ResponseHeaders int32 `json:"response_headers,omitempty"`

	// Number of response_length.
	ResponseLength int64 `json:"response_length,omitempty"`

	// Number of response_time_first_byte.
	ResponseTimeFirstByte int64 `json:"response_time_first_byte,omitempty"`

	// Number of response_time_last_byte.
	ResponseTimeLastByte int64 `json:"response_time_last_byte,omitempty"`

	// rewritten_uri_path of ApplicationLog.
	RewrittenURIPath string `json:"rewritten_uri_path,omitempty"`

	// rewritten_uri_query of ApplicationLog.
	RewrittenURIQuery string `json:"rewritten_uri_query,omitempty"`

	// Number of server_conn_src_ip.
	ServerConnSrcIP int32 `json:"server_conn_src_ip,omitempty"`

	// Flag to indicate if connection from the connection pool was reused.
	ServerConnectionReused bool `json:"server_connection_reused,omitempty"`

	// Number of server_dest_port.
	ServerDestPort int32 `json:"server_dest_port,omitempty"`

	// Number of server_ip.
	ServerIP int32 `json:"server_ip,omitempty"`

	// server_name of ApplicationLog.
	ServerName string `json:"server_name,omitempty"`

	// Number of server_response_code.
	ServerResponseCode int32 `json:"server_response_code,omitempty"`

	// Number of server_response_length.
	ServerResponseLength int64 `json:"server_response_length,omitempty"`

	// Number of server_response_time_first_byte.
	ServerResponseTimeFirstByte int64 `json:"server_response_time_first_byte,omitempty"`

	// Number of server_response_time_last_byte.
	ServerResponseTimeLastByte int64 `json:"server_response_time_last_byte,omitempty"`

	// Number of server_rtt.
	ServerRtt int32 `json:"server_rtt,omitempty"`

	// server_side_redirect_uri of ApplicationLog.
	ServerSideRedirectURI string `json:"server_side_redirect_uri,omitempty"`

	// Number of server_src_port.
	ServerSrcPort int32 `json:"server_src_port,omitempty"`

	// SSL session id for the backend connection.
	ServerSslSessionID string `json:"server_ssl_session_id,omitempty"`

	// Flag to indicate if SSL session was reused.
	ServerSslSessionReused bool `json:"server_ssl_session_reused,omitempty"`

	// service_engine of ApplicationLog.
	// Required: true
	ServiceEngine string `json:"service_engine"`

	// significance of ApplicationLog.
	Significance string `json:"significance,omitempty"`

	// Number of significant.
	// Required: true
	Significant int64 `json:"significant"`

	// List of enums which indicate why a log is significant. Enum options - ADF_CLIENT_CONN_SETUP_REFUSED, ADF_SERVER_CONN_SETUP_REFUSED, ADF_CLIENT_CONN_SETUP_TIMEDOUT, ADF_SERVER_CONN_SETUP_TIMEDOUT, ADF_CLIENT_CONN_SETUP_FAILED_INTERNAL, ADF_SERVER_CONN_SETUP_FAILED_INTERNAL, ADF_CLIENT_CONN_SETUP_FAILED_BAD_PACKET, ADF_UDP_CONN_SETUP_FAILED_INTERNAL, ADF_UDP_SERVER_CONN_SETUP_FAILED_INTERNAL, ADF_CLIENT_SENT_RESET, ADF_SERVER_SENT_RESET, ADF_CLIENT_CONN_TIMEDOUT, ADF_SERVER_CONN_TIMEDOUT, ADF_USER_DELETE_OPERATION, ADF_CLIENT_REQUEST_TIMEOUT, ADF_CLIENT_CONN_ABORTED, ADF_CLIENT_SSL_HANDSHAKE_FAILURE, ADF_CLIENT_CONN_FAILED, ADF_SERVER_CERTIFICATE_VERIFICATION_FAILED, ADF_SERVER_SIDE_SSL_HANDSHAKE_FAILED, ADF_IDLE_TIMEDOUT, ADF_CLIENT_HIGH_TIMEOUT_RETRANSMITS, ADF_SERVER_HIGH_TIMEOUT_RETRANSMITS, ADF_CLIENT_HIGH_RX_ZERO_WINDOW_SIZE_EVENTS, ADF_SERVER_HIGH_RX_ZERO_WINDOW_SIZE_EVENTS, ADF_CLIENT_RTT_ABOVE_SEC, ADF_SERVER_RTT_ABOVE_500MS, ADF_CLIENT_HIGH_TOTAL_RETRANSMITS, ADF_SERVER_HIGH_TOTAL_RETRANSMITS, ADF_CLIENT_HIGH_OUT_OF_ORDERS, ADF_SERVER_HIGH_OUT_OF_ORDERS, ADF_CLIENT_HIGH_TX_ZERO_WINDOW_SIZE_EVENTS, ADF_SERVER_HIGH_TX_ZERO_WINDOW_SIZE_EVENTS, ADF_CLIENT_POSSIBLE_WINDOW_STUCK, ADF_SERVER_POSSIBLE_WINDOW_STUCK, ADF_SERVER_UNANSWERED_SYNS, ADF_RESPONSE_CODE_4XX, ADF_RESPONSE_CODE_5XX, ADF_LOAD_BALANCING_FAILED, ADF_DATASCRIPT_EXECUTION_FAILED, ADF_REQUEST_NO_POOL, ADF_RATE_LIMIT_DROP_CLIENT_IP, ADF_RATE_LIMIT_DROP_URI, ADF_RATE_LIMIT_DROP_CLIENT_IP_URI, ADF_RATE_LIMIT_DROP_UNKNOWN_URI, ADF_RATE_LIMIT_DROP_BAD_URI, ADF_REQUEST_VIRTUAL_HOSTING_APP_SELECT_FAILED, ADF_RATE_LIMIT_DROP_UNKNOWN_CIP, ADF_RATE_LIMIT_DROP_BAD_CIP, ADF_RATE_LIMIT_DROP_CLIENT_IP_BAD, ADF_RATE_LIMIT_DROP_URI_BAD, ADF_RATE_LIMIT_DROP_CLIENT_IP_URI_BAD, ADF_RATE_LIMIT_DROP_REQ, ADF_RATE_LIMIT_DROP_CLIENT_IP_CONN, ADF_RATE_LIMIT_DROP_CONN, ADF_RATE_LIMIT_DROP_HEADER, ADF_HTTP_VERSION_LT_1_0, ADF_CLIENT_HIGH_RESPONSE_TIME, ADF_SERVER_HIGH_RESPONSE_TIME, ADF_PERSISTENT_SERVER_CHANGE, ADF_DOS_SERVER_BAD_GATEWAY, ADF_DOS_SERVER_GATEWAY_TIMEOUT, ADF_DOS_CLIENT_SENT_RESET, ADF_DOS_CLIENT_CONN_TIMEOUT, ADF_DOS_CLIENT_REQUEST_TIMEOUT, ADF_DOS_CLIENT_CONN_ABORTED, ADF_DOS_CLIENT_BAD_REQUEST, ADF_DOS_CLIENT_REQUEST_ENTITY_TOO_LARGE, ADF_DOS_CLIENT_REQUEST_URI_TOO_LARGE, ADF_DOS_CLIENT_REQUEST_HEADER_TOO_LARGE, ADF_DOS_CLIENT_CLOSED_REQUEST, ADF_DOS_SSL_ERROR, ADF_X509_CLIENT_CERTIFICATE_VERIFICATION_FAILED, ADF_X509_CLIENT_CERTIFICATE_NOT_YET_VALID, ADF_X509_CLIENT_CERTIFICATE_EXPIRED, ADF_X509_CLIENT_CERTIFICATE_REVOKED, ADF_X509_CLIENT_CERTIFICATE_INVALID_CA, ADF_X509_CLIENT_CERTIFICATE_CRL_NOT_PRESENT, ADF_X509_CLIENT_CERTIFICATE_CRL_NOT_YET_VALID, ADF_X509_CLIENT_CERTIFICATE_CRL_EXPIRED, ADF_X509_CLIENT_CERTIFICATE_CRL_ERROR, ADF_X509_CLIENT_CERTIFICATE_CHAINING_ERROR, ADF_X509_CLIENT_CERTIFICATE_INTERNAL_ERROR, ADF_X509_CLIENT_CERTIFICATE_FORMAT_ERROR, ADF_UDP_PORT_NOT_REACHABLE, ADF_UDP_CONN_TIMEOUT, ADF_X509_SERVER_CERTIFICATE_VERIFICATION_FAILED, ADF_X509_SERVER_CERTIFICATE_NOT_YET_VALID, ADF_X509_SERVER_CERTIFICATE_EXPIRED, ADF_X509_SERVER_CERTIFICATE_REVOKED, ADF_X509_SERVER_CERTIFICATE_INVALID_CA, ADF_X509_SERVER_CERTIFICATE_CRL_NOT_PRESENT, ADF_X509_SERVER_CERTIFICATE_CRL_NOT_YET_VALID, ADF_X509_SERVER_CERTIFICATE_CRL_EXPIRED, ADF_X509_SERVER_CERTIFICATE_CRL_ERROR, ADF_X509_SERVER_CERTIFICATE_CHAINING_ERROR, ADF_X509_SERVER_CERTIFICATE_INTERNAL_ERROR, ADF_X509_SERVER_CERTIFICATE_FORMAT_ERROR, ADF_X509_SERVER_CERTIFICATE_HOSTNAME_ERROR, ADF_SSL_R_BAD_CHANGE_CIPHER_SPEC, ADF_SSL_R_BLOCK_CIPHER_PAD_IS_WRONG, ADF_SSL_R_DIGEST_CHECK_FAILED, ADF_SSL_R_ERROR_IN_RECEIVED_CIPHER_LIST, ADF_SSL_R_EXCESSIVE_MESSAGE_SIZE, ADF_SSL_R_LENGTH_MISMATCH, ADF_SSL_R_NO_CIPHERS_PASSED, ADF_SSL_R_NO_CIPHERS_SPECIFIED, ADF_SSL_R_NO_COMPRESSION_SPECIFIED, ADF_SSL_R_NO_SHARED_CIPHER, ADF_SSL_R_RECORD_LENGTH_MISMATCH, ADF_SSL_R_PARSE_TLSEXT, ADF_SSL_R_UNEXPECTED_MESSAGE, ADF_SSL_R_UNEXPECTED_RECORD, ADF_SSL_R_UNKNOWN_ALERT_TYPE, ADF_SSL_R_UNKNOWN_PROTOCOL, ADF_SSL_R_WRONG_VERSION_NUMBER, ADF_SSL_R_DECRYPTION_FAILED_OR_BAD_RECORD_MAC, ADF_SSL_R_RENEGOTIATE_EXT_TOO_LONG, ADF_SSL_R_RENEGOTIATION_ENCODING_ERR, ADF_SSL_R_RENEGOTIATION_MISMATCH, ADF_SSL_R_UNSAFE_LEGACY_RENEGOTIATION_DISABLED, ADF_SSL_R_SCSV_RECEIVED_WHEN_RENEGOTIATING, ADF_SSL_R_INAPPROPRIATE_FALLBACK, ADF_SSL_R_SSLV3_ALERT_UNEXPECTED_MESSAGE, ADF_SSL_R_SSLV3_ALERT_BAD_RECORD_MAC, ADF_SSL_R_TLSV1_ALERT_DECRYPTION_FAILED, ADF_SSL_R_TLSV1_ALERT_RECORD_OVERFLOW, ADF_SSL_R_SSLV3_ALERT_DECOMPRESSION_FAILURE, ADF_SSL_R_SSLV3_ALERT_HANDSHAKE_FAILURE, ADF_SSL_R_SSLV3_ALERT_NO_CERTIFICATE, ADF_SSL_R_SSLV3_ALERT_BAD_CERTIFICATE, ADF_SSL_R_SSLV3_ALERT_UNSUPPORTED_CERTIFICATE, ADF_SSL_R_SSLV3_ALERT_CERTIFICATE_REVOKED, ADF_SSL_R_SSLV3_ALERT_CERTIFICATE_EXPIRED, ADF_SSL_R_SSLV3_ALERT_CERTIFICATE_UNKNOWN, ADF_SSL_R_SSLV3_ALERT_ILLEGAL_PARAMETER, ADF_SSL_R_TLSV1_ALERT_UNKNOWN_CA, ADF_SSL_R_TLSV1_ALERT_ACCESS_DENIED, ADF_SSL_R_TLSV1_ALERT_DECODE_ERROR, ADF_SSL_R_TLSV1_ALERT_DECRYPT_ERROR, ADF_SSL_R_TLSV1_ALERT_EXPORT_RESTRICTION, ADF_SSL_R_TLSV1_ALERT_PROTOCOL_VERSION, ADF_SSL_R_TLSV1_ALERT_INSUFFICIENT_SECURITY, ADF_SSL_R_TLSV1_ALERT_INTERNAL_ERROR, ADF_SSL_R_TLSV1_ALERT_USER_CANCELLED, ADF_SSL_R_TLSV1_ALERT_NO_RENEGOTIATION, ADF_CLIENT_AUTH_UNKNOWN_USER, ADF_CLIENT_AUTH_LOGIN_FAILED, ADF_CLIENT_AUTH_MISSING_CREDENTIALS, ADF_CLIENT_AUTH_SERVER_CONN_ERROR, ADF_CLIENT_AUTH_USER_NOT_AUTHORIZED, ADF_CLIENT_AUTH_TIMED_OUT, ADF_CLIENT_AUTH_UNKNOWN_ERROR, ADF_CLIENT_DNS_FAILED_INVALID_QUERY, ADF_CLIENT_DNS_FAILED_INVALID_DOMAIN, ADF_CLIENT_DNS_FAILED_NO_SERVICE, ADF_CLIENT_DNS_FAILED_GS_DOWN, ADF_CLIENT_DNS_FAILED_NO_VALID_GS_MEMBER, ADF_SERVER_DNS_ERROR_RESPONSE, ADF_CLIENT_DNS_FAILED_UNSUPPORTED_QUERY, ADF_MEMORY_EXHAUSTED.
	SignificantLog []string `json:"significant_log,omitempty"`

	// spdy_version of ApplicationLog.
	SpdyVersion string `json:"spdy_version,omitempty"`

	// ssl_cipher of ApplicationLog.
	SslCipher string `json:"ssl_cipher,omitempty"`

	// ssl_session_id of ApplicationLog.
	SslSessionID string `json:"ssl_session_id,omitempty"`

	// ssl_version of ApplicationLog.
	SslVersion string `json:"ssl_version,omitempty"`

	// Number of total_time.
	TotalTime int64 `json:"total_time,omitempty"`

	// Placeholder for description of property udf of obj type ApplicationLog field type str  type boolean
	// Required: true
	Udf bool `json:"udf"`

	// uri_path of ApplicationLog.
	URIPath string `json:"uri_path,omitempty"`

	// uri_query of ApplicationLog.
	URIQuery string `json:"uri_query,omitempty"`

	// user_agent of ApplicationLog.
	UserAgent string `json:"user_agent,omitempty"`

	// user_id of ApplicationLog.
	UserID string `json:"user_id,omitempty"`

	// Number of vcpu_id.
	// Required: true
	VcpuID int32 `json:"vcpu_id"`

	// virtualservice of ApplicationLog.
	// Required: true
	Virtualservice string `json:"virtualservice"`

	//  Field introduced in 17.1.1.
	VsIP int32 `json:"vs_ip,omitempty"`

	// xff of ApplicationLog.
	Xff string `json:"xff,omitempty"`
}
