package dtc

// Sierrachart DTC types as per:
// https://dtcprotocol.org/DTC_Files/DTCProtocol.proto

type DTCVersion int32

const (
	DTCVersion_DTC_VERSION_UNSET DTCVersion = 0
	DTCVersion_CURRENT_VERSION   DTCVersion = 8
)

type DTCMessageType int32

const (
	DTCMessageType_MESSAGE_TYPE_UNSET DTCMessageType = 0
	// Authentication and connection monitoring
	DTCMessageType_LOGON_REQUEST     DTCMessageType = 1
	DTCMessageType_LOGON_RESPONSE    DTCMessageType = 2
	DTCMessageType_HEARTBEAT         DTCMessageType = 3
	DTCMessageType_LOGOFF            DTCMessageType = 5
	DTCMessageType_ENCODING_REQUEST  DTCMessageType = 6
	DTCMessageType_ENCODING_RESPONSE DTCMessageType = 7
	// Market data
	DTCMessageType_MARKET_DATA_REQUEST                                 DTCMessageType = 101
	DTCMessageType_MARKET_DATA_REJECT                                  DTCMessageType = 103
	DTCMessageType_MARKET_DATA_SNAPSHOT                                DTCMessageType = 104
	DTCMessageType_MARKET_DATA_UPDATE_TRADE                            DTCMessageType = 107
	DTCMessageType_MARKET_DATA_UPDATE_TRADE_COMPACT                    DTCMessageType = 112
	DTCMessageType_MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT              DTCMessageType = 134
	DTCMessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR   DTCMessageType = 137
	DTCMessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2 DTCMessageType = 146
	DTCMessageType_MARKET_DATA_UPDATE_TRADE_NO_TIMESTAMP               DTCMessageType = 142
	DTCMessageType_MARKET_DATA_UPDATE_BID_ASK                          DTCMessageType = 108
	DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_COMPACT                  DTCMessageType = 117
	DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP             DTCMessageType = 143
	DTCMessageType_MARKET_DATA_UPDATE_BID_ASK_FLOAT_WITH_MICROSECONDS  DTCMessageType = 144
	DTCMessageType_MARKET_DATA_UPDATE_SESSION_OPEN                     DTCMessageType = 120
	DTCMessageType_MARKET_DATA_UPDATE_SESSION_HIGH                     DTCMessageType = 114
	DTCMessageType_MARKET_DATA_UPDATE_SESSION_LOW                      DTCMessageType = 115
	DTCMessageType_MARKET_DATA_UPDATE_SESSION_VOLUME                   DTCMessageType = 113
	DTCMessageType_MARKET_DATA_UPDATE_OPEN_INTEREST                    DTCMessageType = 124
	DTCMessageType_MARKET_DATA_UPDATE_SESSION_SETTLEMENT               DTCMessageType = 119
	DTCMessageType_MARKET_DATA_UPDATE_SESSION_NUM_TRADES               DTCMessageType = 135
	DTCMessageType_MARKET_DATA_UPDATE_TRADING_SESSION_DATE             DTCMessageType = 136
	DTCMessageType_MARKET_DEPTH_REQUEST                                DTCMessageType = 102
	DTCMessageType_MARKET_DEPTH_REJECT                                 DTCMessageType = 121
	DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL                         DTCMessageType = 122
	DTCMessageType_MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT                   DTCMessageType = 145
	DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL                           DTCMessageType = 106
	DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS   DTCMessageType = 140
	DTCMessageType_MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP              DTCMessageType = 141
	DTCMessageType_MARKET_DATA_FEED_STATUS                             DTCMessageType = 100
	DTCMessageType_MARKET_DATA_FEED_SYMBOL_STATUS                      DTCMessageType = 116
	DTCMessageType_TRADING_SYMBOL_STATUS                               DTCMessageType = 138
	DTCMessageType_MARKET_ORDERS_REQUEST                               DTCMessageType = 150
	DTCMessageType_MARKET_ORDERS_REJECT                                DTCMessageType = 151
	DTCMessageType_MARKET_ORDERS_ADD                                   DTCMessageType = 152
	DTCMessageType_MARKET_ORDERS_MODIFY                                DTCMessageType = 153
	DTCMessageType_MARKET_ORDERS_REMOVE                                DTCMessageType = 154
	DTCMessageType_MARKET_ORDERS_SNAPSHOT_MESSAGE_BOUNDARY             DTCMessageType = 155
	// Order entry and modification
	DTCMessageType_SUBMIT_NEW_SINGLE_ORDER             DTCMessageType = 208
	DTCMessageType_SUBMIT_NEW_OCO_ORDER                DTCMessageType = 201
	DTCMessageType_SUBMIT_FLATTEN_POSITION_ORDER       DTCMessageType = 209
	DTCMessageType_FLATTEN_POSITIONS_FOR_TRADE_ACCOUNT DTCMessageType = 210
	DTCMessageType_CANCEL_ORDER                        DTCMessageType = 203
	DTCMessageType_CANCEL_REPLACE_ORDER                DTCMessageType = 204
	// Trading related
	DTCMessageType_OPEN_ORDERS_REQUEST            DTCMessageType = 300
	DTCMessageType_OPEN_ORDERS_REJECT             DTCMessageType = 302
	DTCMessageType_ORDER_UPDATE                   DTCMessageType = 301
	DTCMessageType_HISTORICAL_ORDER_FILLS_REQUEST DTCMessageType = 303
	DTCMessageType_HISTORICAL_ORDER_FILL_RESPONSE DTCMessageType = 304
	DTCMessageType_HISTORICAL_ORDER_FILLS_REJECT  DTCMessageType = 308
	DTCMessageType_CURRENT_POSITIONS_REQUEST      DTCMessageType = 305
	DTCMessageType_CURRENT_POSITIONS_REJECT       DTCMessageType = 307
	DTCMessageType_POSITION_UPDATE                DTCMessageType = 306
	DTCMessageType_ADD_CORRECTING_ORDER_FILL      DTCMessageType = 309
	DTCMessageType_CORRECTING_ORDER_FILL_RESPONSE DTCMessageType = 310
	// Account list
	DTCMessageType_TRADE_ACCOUNTS_REQUEST DTCMessageType = 400
	DTCMessageType_TRADE_ACCOUNT_RESPONSE DTCMessageType = 401
	// Symbol discovery and security definitions
	DTCMessageType_EXCHANGE_LIST_REQUEST                   DTCMessageType = 500
	DTCMessageType_EXCHANGE_LIST_RESPONSE                  DTCMessageType = 501
	DTCMessageType_SYMBOLS_FOR_EXCHANGE_REQUEST            DTCMessageType = 502
	DTCMessageType_UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST DTCMessageType = 503
	DTCMessageType_SYMBOLS_FOR_UNDERLYING_REQUEST          DTCMessageType = 504
	DTCMessageType_SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  DTCMessageType = 506
	DTCMessageType_SECURITY_DEFINITION_RESPONSE            DTCMessageType = 507
	DTCMessageType_SYMBOL_SEARCH_REQUEST                   DTCMessageType = 508
	DTCMessageType_SECURITY_DEFINITION_REJECT              DTCMessageType = 509
	// Account balance
	DTCMessageType_ACCOUNT_BALANCE_REQUEST             DTCMessageType = 601
	DTCMessageType_ACCOUNT_BALANCE_REJECT              DTCMessageType = 602
	DTCMessageType_ACCOUNT_BALANCE_UPDATE              DTCMessageType = 600
	DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT          DTCMessageType = 607
	DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT_REJECT   DTCMessageType = 608
	DTCMessageType_ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE DTCMessageType = 609
	DTCMessageType_HISTORICAL_ACCOUNT_BALANCES_REQUEST DTCMessageType = 603
	DTCMessageType_HISTORICAL_ACCOUNT_BALANCES_REJECT  DTCMessageType = 604
	DTCMessageType_HISTORICAL_ACCOUNT_BALANCE_RESPONSE DTCMessageType = 606
	// Logging
	DTCMessageType_USER_MESSAGE            DTCMessageType = 700
	DTCMessageType_GENERAL_LOG_MESSAGE     DTCMessageType = 701
	DTCMessageType_ALERT_MESSAGE           DTCMessageType = 702
	DTCMessageType_JOURNAL_ENTRY_ADD       DTCMessageType = 703
	DTCMessageType_JOURNAL_ENTRIES_REQUEST DTCMessageType = 704
	DTCMessageType_JOURNAL_ENTRIES_REJECT  DTCMessageType = 705
	DTCMessageType_JOURNAL_ENTRY_RESPONSE  DTCMessageType = 706
	// Historical price data
	DTCMessageType_HISTORICAL_PRICE_DATA_REQUEST              DTCMessageType = 800
	DTCMessageType_HISTORICAL_PRICE_DATA_RESPONSE_HEADER      DTCMessageType = 801
	DTCMessageType_HISTORICAL_PRICE_DATA_REJECT               DTCMessageType = 802
	DTCMessageType_HISTORICAL_PRICE_DATA_RECORD_RESPONSE      DTCMessageType = 803
	DTCMessageType_HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE DTCMessageType = 804
	DTCMessageType_HISTORICAL_PRICE_DATA_RESPONSE_TRAILER     DTCMessageType = 807
	// Historical market depth data
	DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_REQUEST         DTCMessageType = 900
	DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_RESPONSE_HEADER DTCMessageType = 901
	DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_REJECT          DTCMessageType = 902
	DTCMessageType_HISTORICAL_MARKET_DEPTH_DATA_RECORD_RESPONSE DTCMessageType = 903
	// Nonstandard
	DTCMessageType_TRADE_ACCOUNT_TRADING_IS_DISABLED_REQUEST  DTCMessageType = 10206
	DTCMessageType_TRADE_ACCOUNT_TRADING_IS_DISABLED_RESPONSE DTCMessageType = 10207
	DTCMessageType_TRADE_ACCOUNT_DATA_DUPLICATE               DTCMessageType = 10208
)

type OpenCloseTradeEnum int32

const (
	OpenCloseTradeEnum_TRADE_UNSET OpenCloseTradeEnum = 0
	OpenCloseTradeEnum_TRADE_OPEN  OpenCloseTradeEnum = 1
	OpenCloseTradeEnum_TRADE_CLOSE OpenCloseTradeEnum = 2
)

type BuySellEnum int32

const (
	BuySellEnum_BUY_SELL_UNSET BuySellEnum = 0
	BuySellEnum_BUY            BuySellEnum = 1
	BuySellEnum_SELL           BuySellEnum = 2
)

type HistoricalOrderFillResponse struct {
	Type                    DTCMessageType
	RequestID               int32              `json:"RequestID,omitempty"`
	TotalNumberMessages     int32              `json:"TotalNumberMessages,omitempty"`
	MessageNumber           int32              `json:"MessageNumber,omitempty"`
	Symbol                  string             `json:"Symbol,omitempty"`
	Exchange                string             `json:"Exchange,omitempty"`
	ServerOrderID           string             `json:"ServerOrderID,omitempty"`
	BuySell                 BuySellEnum        `json:"BuySell,omitempty"`
	Price                   float64            `json:"Price,omitempty"`
	DateTime                int64              `json:"DateTime,omitempty"`
	Quantity                float64            `json:"Quantity,omitempty"`
	UniqueExecutionID       string             `json:"UniqueExecutionID,omitempty"`
	TradeAccount            string             `json:"TradeAccount,omitempty"`
	OpenClose               OpenCloseTradeEnum `json:"OpenClose,omitempty"`
	NoOrderFills            uint32             `json:"NoOrderFills,omitempty"`
	InfoText                string             `json:"InfoText,omitempty"`
	HighPriceDuringPosition float64            `json:"HighPriceDuringPosition,omitempty"`
	LowPriceDuringPosition  float64            `json:"LowPriceDuringPosition,omitempty"`
	PositionQuantity        float64            `json:"PositionQuantity,omitempty"`
	Username                string             `json:"Username,omitempty"`
	ExchangeOrderID         string             `json:"ExchangeOrderID,omitempty"`
	SenderSubID             string             `json:"SenderSubID,omitempty"`
}

type FillsRequest struct {
	Type          DTCMessageType
	RequestID     int
	ServerOrderID int
	TradeAccount  string
	NumberOfDays  int
	StartDateTime int
}

type LogonRequest struct {
	Type                           DTCMessageType
	ProtocolVersion                int32
	Username                       string
	Password                       string
	GeneralTextData                string
	Integer_1                      int32
	Integer_2                      int32
	HeartbeatIntervalInSeconds     int32
	Unused1                        int32
	TradeAccount                   string
	HardwareIdentifier             string
	ClientName                     string
	MarketDataTransmissionInterval int32
}

type LogonStatusEnum int32

const (
	LogonStatusEnum_LOGON_STATUS_UNSET          LogonStatusEnum = 0
	LogonStatusEnum_LOGON_SUCCESS               LogonStatusEnum = 1
	LogonStatusEnum_LOGON_ERROR                 LogonStatusEnum = 2
	LogonStatusEnum_LOGON_ERROR_NO_RECONNECT    LogonStatusEnum = 3
	LogonStatusEnum_LOGON_RECONNECT_NEW_ADDRESS LogonStatusEnum = 4
)

type LogonResponse struct {
	Type                                          DTCMessageType
	ProtocolVersion                               int32           `json:"ProtocolVersion,omitempty"`
	Result                                        LogonStatusEnum `json:"Result,omitempty"`
	ResultText                                    string          `json:"ResultText,omitempty"`
	ReconnectAddress                              string          `json:"ReconnectAddress,omitempty"`
	Integer_1                                     int32           `json:"Integer_1,omitempty"`
	ServerName                                    string          `json:"ServerName,omitempty"`
	MarketDepthUpdatesBestBidAndAsk               uint32          `json:"MarketDepthUpdatesBestBidAndAsk,omitempty"`
	TradingIsSupported                            uint32          `json:"TradingIsSupported,omitempty"`
	OCOOrdersSupported                            uint32          `json:"OCOOrdersSupported,omitempty"`
	OrderCancelReplaceSupported                   uint32          `json:"OrderCancelReplaceSupported,omitempty"`
	SymbolExchangeDelimiter                       string          `json:"SymbolExchangeDelimiter,omitempty"`
	SecurityDefinitionsSupported                  uint32          `json:"SecurityDefinitionsSupported,omitempty"`
	HistoricalPriceDataSupported                  uint32          `json:"HistoricalPriceDataSupported,omitempty"`
	ResubscribeWhenMarketDataFeedAvailable        uint32          `json:"ResubscribeWhenMarketDataFeedAvailable,omitempty"`
	MarketDepthIsSupported                        uint32          `json:"MarketDepthIsSupported,omitempty"`
	OneHistoricalPriceDataRequestPerConnection    uint32          `json:"OneHistoricalPriceDataRequestPerConnection,omitempty"`
	BracketOrdersSupported                        uint32          `json:"BracketOrdersSupported,omitempty"`
	Unused_1                                      uint32          `json:"Unused_1,omitempty"`
	UsesMultiplePositionsPerSymbolAndTradeAccount uint32          `json:"UsesMultiplePositionsPerSymbolAndTradeAccount,omitempty"`
	MarketDataSupported                           uint32          `json:"MarketDataSupported,omitempty"`
}

type Heartbeat struct {
	Type               DTCMessageType
	NumDroppedMessages uint32
	CurrentDateTime    int64
}

type MessageBase struct {
	Type DTCMessageType
}
