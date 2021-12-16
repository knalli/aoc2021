package day16

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BitsMessage struct {
	Version int
	Type    int

	Consumed int

	// type=4
	Value int

	// type!=4 (op)
	LengthId    int
	SubMessages []BitsMessage
}

func (m *BitsMessage) ToString() string {
	res := fmt.Sprintf("V=%d T=%d c=%d", m.Version, m.Type, m.Consumed)
	if m.Type == 4 {
		res += fmt.Sprintf(" [value=%d]", m.Value)
	} else {
		res += " {\n"
		for _, s := range m.SubMessages {
			for _, line := range strings.Split(s.ToString(), "\n") {
				res += "  " + line + "\n"
			}
		}
		res += "}"
	}
	return res
}

func parseBitsMessage(input string) (*BitsMessage, error) {
	messageConsumed := 0
	version, err := parseBinToInt(input[0:3])
	input = input[3:]
	messageConsumed += 3
	if err != nil {
		return nil, err
	}
	typeId, err := parseBinToInt(input[0:3])
	input = input[3:]
	messageConsumed += 3
	if err != nil {
		return nil, err
	}

	if typeId == 4 {
		val, consumed, err := parseLiteralValuePayload(input)
		input = input[consumed:]
		messageConsumed += consumed
		if err != nil {
			return nil, err
		}
		return &BitsMessage{
			Version:  version,
			Type:     typeId,
			Value:    val,
			Consumed: messageConsumed,
		}, nil

	} else {
		lengthTypeId, err := parseBinToInt(input[0:1])
		input = input[1:]
		messageConsumed += 1
		if err != nil {
			return nil, err
		}
		if lengthTypeId == 0 {
			totalLength, err := parseBinToInt(input[0:15])
			if err != nil {
				return nil, err
			}
			input = input[15:]
			messageConsumed += 15
			subMessages := make([]BitsMessage, 0)
			for consumable := totalLength; consumable > 0; {
				subMessage, err := parseBitsMessage(input)
				if err != nil {
					return nil, err
				}
				subMessages = append(subMessages, *subMessage)
				input = input[subMessage.Consumed:]
				messageConsumed += subMessage.Consumed
				consumable -= subMessage.Consumed
			}
			return &BitsMessage{
				Version:     version,
				Type:        typeId,
				SubMessages: subMessages,
				Consumed:    messageConsumed,
			}, nil

		} else if lengthTypeId == 1 {
			totalPackets, err := parseBinToInt(input[0:11])
			if err != nil {
				return nil, err
			}
			input = input[11:]
			messageConsumed += 11
			subMessages := make([]BitsMessage, 0)
			for i := 0; i < totalPackets; i++ {
				subMessage, err := parseBitsMessage(input)
				if err != nil {
					return nil, err
				}
				subMessages = append(subMessages, *subMessage)
				input = input[subMessage.Consumed:]
				messageConsumed += subMessage.Consumed
			}
			return &BitsMessage{
				Version:     version,
				Type:        typeId,
				SubMessages: subMessages,
				Consumed:    messageConsumed,
			}, nil

		} else {
			return nil, errors.New("invalid length type id")
		}
	}

	return nil, errors.New("unknown message format")
}

func parseBinToInt(str string) (int, error) {
	res, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		return -1, err
	} else {
		return int(res), nil
	}
}

func parseLiteralValuePayload(input string) (int, int, error) {
	res := ""
	consumed := 0
	for i := 0; i < len(input); i += 5 {
		res += input[i+1 : i+5]
		consumed += 5
		if input[i] == '0' {
			break
		}
	}
	val, err := parseBinToInt(res)
	if err != nil {
		return -1, -1, err
	}
	return val, consumed, nil
}
