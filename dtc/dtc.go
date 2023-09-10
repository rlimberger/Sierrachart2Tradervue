package dtc

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type Client struct {
	conn               *net.Conn
	fillRequestHandler *func(HistoricalOrderFillResponse)
}

func NewClient() (*Client, error) {

	client := Client{}

	conn, err := net.Dial("tcp", "127.0.0.1:11099")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create connection to DTC server %s", err.Error()))
	}

	//defer conn.Close()
	client.conn = &conn

	err = client.doLogon()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed logon to DTC server %s", err.Error()))
	}

	go client.doHeartbeat()

	log.Println("Successfully logged on to DTC server")

	return &client, nil
}

func (c *Client) doLogon() error {
	request := LogonRequest{
		Type:                           DTCMessageType_LOGON_REQUEST,
		ProtocolVersion:                int32(DTCVersion_CURRENT_VERSION),
		Username:                       "",
		Password:                       "",
		GeneralTextData:                "",
		Integer_1:                      0,
		Integer_2:                      0,
		HeartbeatIntervalInSeconds:     60,
		Unused1:                        0,
		TradeAccount:                   "",
		HardwareIdentifier:             "",
		ClientName:                     "Sierrachart2Tradervue",
		MarketDataTransmissionInterval: 0,
	}

	err := c.sendRequest(request)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to send request %s", err.Error()))
	}

	var response LogonResponse
	r := bufio.NewReader(*c.conn)
	for {
		line, err := r.ReadBytes(byte(0))
		switch err {
		case nil:
			line = line[:len(line)-1] // remove null terminator
			err := json.Unmarshal(line, &response)
			if err != nil {
				return errors.New(fmt.Sprintf("error reading LogonResponse %s", err.Error()))
			}
			return nil
		case io.EOF:
		default:
			return errors.New(fmt.Sprintf("error reading LogonResponse %s", err.Error()))
		}
	}

	return errors.New("unable to read of LogonResponse")
}

func (c *Client) doHeartbeat() {
	for {
		time.Sleep(5 + time.Second)
		hb := Heartbeat{
			Type:               DTCMessageType_HEARTBEAT,
			NumDroppedMessages: 0,
			CurrentDateTime:    time.Now().Unix(),
		}
		_ = c.sendRequest(hb)
	}
}

func (c *Client) sendRequest(data any) error {
	bytes, err := json.Marshal(&data)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to encode request data %s", err.Error()))
	}

	// null terminate
	bytes = append(bytes, 0)

	// write to socket
	written, err := (*c.conn).Write(bytes)
	if err != nil || written != len(bytes) {
		return errors.New(fmt.Sprintf("failed to write Request %s", err.Error()))
	}

	return nil
}

func (c *Client) RequestHistoricalFills(tradeAccount string, numberOfDays int) ([]HistoricalOrderFillResponse, error) {
	request := FillsRequest{
		Type:          DTCMessageType_HISTORICAL_ORDER_FILLS_REQUEST,
		RequestID:     1,
		ServerOrderID: 0,
		TradeAccount:  tradeAccount,
		NumberOfDays:  numberOfDays,
		StartDateTime: 0,
	}

	err := c.sendRequest(request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to send request %s", err.Error()))
	}

	responses := make([]HistoricalOrderFillResponse, 0)
	var response HistoricalOrderFillResponse
	r := bufio.NewReader(*c.conn)
	for {
		line, err := r.ReadBytes(byte(0))
		switch err {
		case nil:
			line = line[:len(line)-1] // remove null terminator
			err := json.Unmarshal(line, &response)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("error reading HistoricalOrderFillResponse %s", err.Error()))
			}
			if response.Type != DTCMessageType_HISTORICAL_ORDER_FILL_RESPONSE {
				continue
			}
			responses = append(responses, response)

			// check if this is the last one
			if response.TotalNumberMessages == response.MessageNumber {
				return responses, nil
			}

			break
		case io.EOF:
		default:
			return nil, errors.New(fmt.Sprintf("error reading HistoricalOrderFillResponse %s", err.Error()))
		}
	}

	return responses, errors.New("incomplete read of HistoricalFill responses")
}
