package rabbitmq

// ProcessCampaign ...
type ProcessCampaign struct {
	QueueName string `json:"queue_name"`
	ActorName string `json:"actor_name"`
	Args      []Arg  `json:"args"`
	Kwargs    struct {
	} `json:"kwargs"`
	Options struct {
	} `json:"options"`
	MessageID        string `json:"message_id"`
	MessageTimestamp int64  `json:"message_timestamp"`
}

// Arg ...
type Arg struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Dsp  struct {
		ID        int      `json:"id"`
		Name      string   `json:"name"`
		BidURL    string   `json:"bid_url"`
		Active    bool     `json:"active"`
		ExtFields []string `json:"ext_fields"`
	} `json:"dsp,omitempty"`
	SubscriberSelectionSize int         `json:"subscriber_selection_size,omitempty"`
	Targetings              []Targeting `json:"targetings,omitempty"`
	BidsVolume              int         `json:"bids_volume,omitempty"`
	PushLimitPerToken       int         `json:"push_limit_per_token,omitempty"`
	StartHour               int         `json:"start_hour,omitempty"`
	EndHour                 int         `json:"end_hour,omitempty"`
	TokenBidInterval        int         `json:"token_bid_interval,omitempty"`
}

// Targeting ...
type Targeting struct {
	Field    string   `json:"field"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

// Subscriber ...
type Subscriber struct {
	MetaID         string      `json:"_id"`
	PushSenderID   string      `json:"push_sender_id"`
	CountryCode    string      `json:"country_code"`
	Date           string      `json:"date"`
	Model          string      `json:"model"`
	Browser        string      `json:"browser"`
	SiteID         string      `json:"site_id"`
	SafeUID        string      `json:"safe_uid"`
	Subacc         string      `json:"subacc"`
	Bversion       string      `json:"bversion"`
	Subacc2        string      `json:"subacc2"`
	Os             string      `json:"os"`
	WvMinor        string      `json:"wv_minor"`
	Osversion      string      `json:"osversion"`
	SlavePrefixID  string      `json:"slave_prefix_id"`
	BrowserLocale  string      `json:"browser_locale"`
	LastOrder      interface{} `json:"last_order"`
	DeviceType     string      `json:"device_type"`
	OfferID        string      `json:"offer_id"`
	Esub           string      `json:"esub"`
	ID             string      `json:"id"`
	Apsubid3       string      `json:"apsubid3"`
	IabCat         string      `json:"iab_cat"`
	UserSafeID     interface{} `json:"user_safe_id"`
	SiteOption     string      `json:"site_option"`
	Timezone       string      `json:"timezone"`
	PrelandID      string      `json:"preland_id"`
	UID            string      `json:"uid"`
	UserAgent      string      `json:"user_agent"`
	FromURL        string      `json:"from_url"`
	SiteOptionID   string      `json:"site_option_id"`
	Subtool        string      `json:"subtool"`
	Subacc4        string      `json:"subacc4"`
	Etag           string      `json:"etag"`
	Unsub          string      `json:"unsub"`
	RegDate        string      `json:"reg_date"`
	Ap             string      `json:"ap"`
	GeoCountryName string      `json:"geo_country_name"`
	GeoContinent   string      `json:"geo_continent"`
	IP             string      `json:"ip"`
	Subacc3        string      `json:"subacc3"`
	Useragent      string      `json:"useragent"`
	Brand          string      `json:"brand"`
	Node           string      `json:"node"`
	Failed         string      `json:"failed"`
	GeoCountry     string      `json:"geo_country"`
	WvMajor        string      `json:"wv_major"`
	City           string      `json:"city"`
	Countryname    string      `json:"countryname"`
	UserOrigID     string      `json:"user_orig_id"`
	OrderPlaced    string      `json:"order_placed"`
	IPAddress      string      `json:"ip_address"`
	Platform       string      `json:"platform"`
	Country        string      `json:"country"`
	Device         string      `json:"device"`
	Sex            string      `json:"sex"`
	GeoCity        string      `json:"geo_city"`
	Eproot         string      `json:"eproot"`
	Site           string      `json:"site"`
	Browserversion string      `json:"browserversion"`
	ServerIP       string      `json:"server_ip"`
}
