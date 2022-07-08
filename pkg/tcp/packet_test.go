package server

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_pack_unpack(t *testing.T) {
	t.Run("when everything's fine", func(t *testing.T) {
		packer := NewDefaultPacker()
		msg := NewMessage(1, []byte{1, 2, 3})
		data := packer.Pack(msg)
		buffer := bytes.NewBuffer(data)
		msg2, err := packer.Unpack(buffer)
		assert.NoError(t, err)
		assert.Equal(t, msg, msg2)
	})
}

func Test_mock_packer(t *testing.T) {
	t.Run("when everything's fine", func(t *testing.T) {
		ctl := gomock.NewController(t)
		mockPacker := NewMockPacker(ctl)
		mockPacker.EXPECT().Pack(gomock.Any()).Return([]byte{1, 2, 3}).AnyTimes()
		pack := mockPacker.Pack(&Message{Id: 1, Data: []byte{1}})
		log.Println(pack)
	})
}
