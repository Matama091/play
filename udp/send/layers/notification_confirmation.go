package layers

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type NotificationConfirmation struct {
	BaseHeader BaseHeader
	PathID     ID
	HMAC       HMAC
}

const (
	baseHeaderLength = 32
	HMACLength       = 16
)

func (noticonf *NotificationConfirmation) GenerateNotificationConfirmation(PathID, ID []byte) {
	SerializeVersion(&noticonf.BaseHeader)
	SerializeType(&noticonf.BaseHeader, TypeClassNotificationConfirmation)
	noticonf.BaseHeader.ID = ID
	noticonf.BaseHeader.TransactionID, _ = GenerateTransactionID()
	noticonf.BaseHeader.SequenceNumber, _ = GenerateSequenceNumber()

	noticonf.PathID = PathID

	fmt.Println("Generate Notification Confirmation", "Notification Confirmation", noticonf)
}

func (noticonf *NotificationConfirmation) Marshal(key []byte) ([]byte, error) {
	planeBuf := make([]byte, 0, noticonf.BaseHeader.MessageLength-baseHeaderLength-HMACLength)
	planeBuffer := bytes.NewBuffer(planeBuf)
	if err := binary.Write(planeBuffer, binary.BigEndian, noticonf.PathID); err != nil {
		return nil, err
	}

	chiperBuf := EncryptPacket(planeBuffer.Bytes(), key)
	noticonf.BaseHeader.MessageLength = BaseHeaderLen + uint16(len(chiperBuf)) + HMACLen

	buf := make([]byte, 0, noticonf.BaseHeader.MessageLength)
	buffer := bytes.NewBuffer(buf)

	if err := binary.Write(buffer, binary.BigEndian, noticonf.BaseHeader.TransactionID); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.BigEndian, noticonf.BaseHeader.Version); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.BigEndian, noticonf.BaseHeader.Flag); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.BigEndian, noticonf.BaseHeader.Type); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.BigEndian, noticonf.BaseHeader.Count); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.BigEndian, noticonf.BaseHeader.SequenceNumber); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.BigEndian, noticonf.BaseHeader.MessageLength); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.BigEndian, noticonf.BaseHeader.NextOpt); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.BigEndian, noticonf.BaseHeader.Reserved); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.BigEndian, noticonf.BaseHeader.ID); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.BigEndian, chiperBuf); err != nil {
		return nil, err
	}

	noticonf.HMAC = GenerateHMAC(buffer.Bytes(), GenerateSalt(&noticonf.BaseHeader))

	if err := binary.Write(buffer, binary.BigEndian, noticonf.HMAC); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
