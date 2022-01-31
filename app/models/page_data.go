package models

type PageData struct {
	Msg  Message
	Data RawData
}

type PageDataJSON struct {
	Msgs []MessageJSON
	Data RawData
}
