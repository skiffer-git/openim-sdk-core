//go:build js && wasm
// +build js,wasm

package indexdb

import "context"

import (
	"open_im_sdk/pkg/db/model_struct"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/sdk_struct"
	"open_im_sdk/wasm/indexdb/temp_struct"
)

type LocalSuperGroupChatLogs struct{}

func NewLocalSuperGroupChatLogs() *LocalSuperGroupChatLogs {
	return &LocalSuperGroupChatLogs{}
}

func (i *LocalSuperGroupChatLogs) GetSuperGroupNormalMsgSeq(ctx context.Context, groupID string) (uint32, error) {
	seq, err := Exec(groupID)
	if err != nil {
		return 0, err
	} else {
		if v, ok := seq.(float64); ok {
			var result uint32
			result = uint32(v)
			return result, err
		} else {
			return 0, ErrType
		}
	}
}
func (i *LocalSuperGroupChatLogs) SuperGroupGetNormalMinSeq(ctx context.Context, groupID string) (uint32, error) {
	seq, err := Exec(groupID)
	if err != nil {
		return 0, err
	} else {
		if v, ok := seq.(float64); ok {
			var result uint32
			result = uint32(v)
			return result, err
		} else {
			return 0, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupGetMessage(ctx context.Context, message *sdk_struct.MsgStruct) (*model_struct.LocalChatLog, error) {
	msg, err := Exec(message.GroupID, message.ClientMsgID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msg.(string); ok {
			result := model_struct.LocalChatLog{}
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return &result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupUpdateMessage(ctx context.Context, c *model_struct.LocalChatLog) error {
	if c.ClientMsgID == "" {
		return PrimaryKeyNull
	}
	tempLocalChatLog := temp_struct.LocalChatLog{
		ServerMsgID:          c.ServerMsgID,
		SendID:               c.SendID,
		RecvID:               c.RecvID,
		SenderPlatformID:     c.SenderPlatformID,
		SenderNickname:       c.SenderNickname,
		SenderFaceURL:        c.SenderFaceURL,
		SessionType:          c.SessionType,
		MsgFrom:              c.MsgFrom,
		ContentType:          c.ContentType,
		Content:              c.Content,
		IsRead:               c.IsRead,
		Status:               c.Status,
		Seq:                  c.Seq,
		SendTime:             c.SendTime,
		CreateTime:           c.CreateTime,
		AttachedInfo:         c.AttachedInfo,
		Ex:                   c.Ex,
		IsReact:              c.IsReact,
		IsExternalExtensions: c.IsExternalExtensions,
		MsgFirstModifyTime:   c.MsgFirstModifyTime,
	}
	_, err := Exec(c.RecvID, c.ClientMsgID, utils.StructToJsonString(tempLocalChatLog))
	return err

}
func (i *LocalSuperGroupChatLogs) SuperGroupBatchInsertMessageList(ctx context.Context, messageList []*model_struct.LocalChatLog, groupID string) error {
	_, err := Exec(utils.StructToJsonString(messageList), groupID)
	return err
}
func (i *LocalSuperGroupChatLogs) SuperGroupInsertMessage(ctx context.Context, message *model_struct.LocalChatLog, groupID string) error {
	_, err := Exec(utils.StructToJsonString(message), groupID)
	return err
}
func (i *LocalSuperGroupChatLogs) SuperGroupGetMultipleMessage(ctx context.Context, msgIDList []string, groupID string) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(utils.StructToJsonString(msgIDList), groupID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}
func (i *LocalSuperGroupChatLogs) SuperGroupUpdateMessageTimeAndStatus(ctx context.Context, msg *sdk_struct.MsgStruct) error {
	_, err := Exec(msg.GroupID, msg.ClientMsgID, msg.ServerMsgID, msg.SendTime, msg.Status)
	return err
}
func (i *LocalSuperGroupChatLogs) SuperGroupGetMessageListNoTime(ctx context.Context, sourceID string, sessionType, count int, isReverse bool) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(sourceID, sessionType, count, isReverse)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}
func (i *LocalSuperGroupChatLogs) SuperGroupGetMessageList(ctx context.Context, sourceID string, sessionType, count int, startTime int64, isReverse bool) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(sourceID, sessionType, count, startTime, isReverse)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupSearchMessageByKeyword(ctx context.Context, contentType []int, keywordList []string, keywordListMatchType int, sourceID string, startTime, endTime int64, sessionType, offset, count int) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(utils.StructToJsonString(contentType), utils.StructToJsonString(keywordList), keywordListMatchType, sourceID, startTime, endTime, sessionType, offset, count)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}
func (i *LocalSuperGroupChatLogs) SuperGroupSearchAllMessageByContentType(ctx context.Context, superGroupID string, contentType int32) (result []*model_struct.LocalChatLog, err error) {
	msgList, err := Exec(superGroupID, contentType)
	if err != nil {
		return nil, err
	} else {
		if v, ok := msgList.(string); ok {
			var temp []*model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) InitSuperLocalErrChatLog(ctx context.Context, groupID string) {
	_, _ = Exec(groupID)
}
func (i *LocalSuperGroupChatLogs) SuperBatchInsertExceptionMsg(ctx context.Context, MessageList []*model_struct.LocalErrChatLog, groupID string) error {
	_, err := Exec(utils.StructToJsonString(MessageList), groupID)
	return err
}

func (i *LocalSuperGroupChatLogs) InitSuperLocalChatLog(ctx context.Context, groupID string) {
	_, _ = Exec(groupID)
}

func (i *LocalSuperGroupChatLogs) SuperGroupDeleteAllMessage(ctx context.Context, groupID string) error {
	_, err := Exec(groupID)
	return err
}

func (i *LocalSuperGroupChatLogs) SuperGroupSearchMessageByContentTypeAndKeyword(ctx context.Context, contentType []int, keywordList []string, keywordListMatchType int, startTime, endTime int64, groupID string) (result []*model_struct.LocalChatLog, err error) {
	gList, err := Exec(utils.StructToJsonString(contentType), utils.StructToJsonString(keywordList), keywordListMatchType, startTime, endTime, groupID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := gList.(string); ok {
			var temp []model_struct.LocalChatLog
			err := utils.JsonStringToStruct(v, &temp)
			if err != nil {
				return nil, err
			}
			for _, v := range temp {
				v1 := v
				result = append(result, &v1)
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupBatchUpdateMessageList(ctx context.Context, MessageList []*model_struct.LocalChatLog) error {
	_, err := Exec(utils.StructToJsonString(MessageList))
	return err
}

func (i *LocalSuperGroupChatLogs) SuperGroupMessageIfExists(ctx context.Context, ClientMsgID string) (bool, error) {
	isExist, err := Exec(ClientMsgID)
	if err != nil {
		return false, err
	} else {
		if v, ok := isExist.(bool); ok {
			return v, nil
		} else {
			return false, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupIsExistsInErrChatLogBySeq(ctx context.Context, seq int64) bool {
	isExist, err := Exec(seq)
	if err != nil {
		return false
	} else {
		if v, ok := isExist.(bool); ok {
			return v
		} else {
			return false
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupMessageIfExistsBySeq(ctx context.Context, seq int64) (bool, error) {
	isExist, err := Exec(seq)
	if err != nil {
		return false, err
	} else {
		if v, ok := isExist.(bool); ok {
			return v, nil
		} else {
			return false, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupGetAllUnDeleteMessageSeqList(ctx context.Context) ([]uint32, error) {
	gList, err := Exec()
	if err != nil {
		return nil, err
	} else {
		if v, ok := gList.(string); ok {
			var result []uint32
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupUpdateColumnsMessage(ctx context.Context, clientMsgID, groupID string, args map[string]interface{}) error {
	_, err := Exec(clientMsgID, groupID, utils.StructToJsonString(args))
	return err
}

func (i *LocalSuperGroupChatLogs) SuperGroupUpdateMessageStatusBySourceID(ctx context.Context, sourceID string, status, sessionType int32) error {
	_, err := Exec(sourceID, status, sessionType)
	return err
}

func (i *LocalSuperGroupChatLogs) SuperGroupGetSendingMessageList(ctx context.Context, groupID string) (result []*model_struct.LocalChatLog, err error) {
	gList, err := Exec(groupID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := gList.(string); ok {
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupUpdateGroupMessageHasRead(ctx context.Context, msgIDList []string, groupID string) error {
	_, err := Exec(utils.StructToJsonString(msgIDList), groupID)
	return err
}

func (i *LocalSuperGroupChatLogs) SuperGroupGetNormalMsgSeq(ctx context.Context) (uint32, error) {
	seq, err := Exec()
	if err != nil {
		return 0, err
	} else {
		if v, ok := seq.(float64); ok {
			return uint32(v), nil
		} else {
			return 0, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupGetTestMessage(ctx context.Context, seq uint32) (*model_struct.LocalChatLog, error) {
	c, err := Exec(seq)
	if err != nil {
		return nil, err
	} else {
		if v, ok := c.(string); ok {
			result := model_struct.LocalChatLog{}
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return &result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupUpdateMsgSenderNickname(ctx context.Context, sendID, nickname string, sType int) error {
	_, err := Exec(sendID, nickname, sType)
	return err
}

func (i *LocalSuperGroupChatLogs) SuperGroupUpdateMsgSenderFaceURL(ctx context.Context, sendID, faceURL string, sType int) error {
	_, err := Exec(sendID, faceURL, sType)
	return err
}

func (i *LocalSuperGroupChatLogs) SuperGroupUpdateMsgSenderFaceURLAndSenderNickname(ctx context.Context, sendID, faceURL, nickname string, sessionType int, groupID string) error {
	_, err := Exec(sendID, faceURL, nickname, sessionType, groupID)
	return err
}

func (i *LocalSuperGroupChatLogs) SuperGroupGetMsgSeqByClientMsgID(ctx context.Context, clientMsgID string, groupID string) (uint32, error) {
	seq, err := Exec(clientMsgID, groupID)
	if err != nil {
		return 0, err
	} else {
		if v, ok := seq.(float64); ok {
			return uint32(v), nil
		} else {
			return 0, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupGetMsgSeqListByGroupID(ctx context.Context, groupID string) ([]uint32, error) {
	gList, err := Exec(groupID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := gList.(string); ok {
			var result []uint32
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupGetMsgSeqListByPeerUserID(ctx context.Context, userID string) ([]uint32, error) {
	gList, err := Exec(userID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := gList.(string); ok {
			var result []uint32
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}

func (i *LocalSuperGroupChatLogs) SuperGroupGetMsgSeqListBySelfUserID(ctx context.Context, userID string) ([]uint32, error) {
	gList, err := Exec(userID)
	if err != nil {
		return nil, err
	} else {
		if v, ok := gList.(string); ok {
			var result []uint32
			err := utils.JsonStringToStruct(v, &result)
			if err != nil {
				return nil, err
			}
			return result, err
		} else {
			return nil, ErrType
		}
	}
}
