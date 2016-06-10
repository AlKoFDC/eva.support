package message

// These are subtypes that a message of TypeMessage can have.
const (
	SubtypeBot     = "bot_message"
	SubtypeMe      = "me_message"
	SubtypeChanged = "message_changed"
	SubtypeDeleted = "message_deleted"

	SubtypeChannelJoin      = "channel_join"
	SubtypeChannelLeave     = "channel_leave"
	SubtypeChannelTopic     = "channel_topic"
	SubtypeChannelPurpose   = "channel_purpose"
	SubtypeChannelName      = "channel_name"
	SubtypeChannelArchive   = "channel_archive"
	SubtypeChannelUnarchive = "channel_unarchive"

	SubtypeGroupJoin      = "group_join"
	SubtypeGroupLeave     = "group_leave"
	SubtypeGroupTopic     = "group_topic"
	SubtypeGroupPurpose   = "group_purpose"
	SubtypeGroupName      = "group_name"
	SubtypeGroupArchive   = "group_archive"
	SubtypeGroupUnarchive = "group_unarchive"

	SubtypeFileShare   = "file_share"
	SubtypeFileComment = "file_comment"
	SubtypeFileMention = "file_mention"

	SubtypePinnedItem   = "pinned_item"
	SubtypeUnpinnedItem = "unpinned_item"
)
