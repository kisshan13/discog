package discog

import "github.com/bwmarrin/discordgo"

// Represents a message embed.
type MessageEmbed struct {
	embed discordgo.MessageEmbed
}

// Represents a message embed auth
type MessageEmbedAuthor struct {
	author discordgo.MessageEmbedAuthor
}

// Returns a MessageEmbed struct
func NewMessageEmbed() *MessageEmbed {
	return &MessageEmbed{
		embed: discordgo.MessageEmbed{},
	}
}

// Sets type for the message embed
func (m *MessageEmbed) SetType(embedType discordgo.EmbedType) *MessageEmbed {
	m.embed.Type = embedType
	return m
}

// Sets Title for the message embed
func (m *MessageEmbed) SetTitle(title string) *MessageEmbed {
	m.embed.Title = title
	return m
}

// Sets Description for the message embed
func (m *MessageEmbed) SetDescription(description string) *MessageEmbed {
	m.embed.Description = description
	return m
}

// Sets Color for the message embed
func (m *MessageEmbed) SetColor(color int) *MessageEmbed {
	m.embed.Color = color
	return m
}

// Sets URL for the message embed
func (m *MessageEmbed) SetURL(url string) *MessageEmbed {
	m.embed.URL = url
	return m
}

// Sets Timestamp for the message embed
func (m *MessageEmbed) SetTimestamp(timestamp string) *MessageEmbed {
	m.embed.Timestamp = timestamp
	return m
}

// Sets Footer for the message embed
func (m *MessageEmbed) SetFooter(footer *discordgo.MessageEmbedFooter) *MessageEmbed {
	m.embed.Footer = footer
	return m
}

// Sets Image for the message embed
func (m *MessageEmbed) SetImage(image *discordgo.MessageEmbedImage) *MessageEmbed {
	m.embed.Image = image
	return m
}

// Sets Thumbnail for the message embed
func (m *MessageEmbed) SetThumbnail(thumbnail *discordgo.MessageEmbedThumbnail) *MessageEmbed {
	m.embed.Thumbnail = thumbnail
	return m
}

// Sets Video for the message embed
func (m *MessageEmbed) SetVideo(video *discordgo.MessageEmbedVideo) *MessageEmbed {
	m.embed.Video = video
	return m
}

// Sets Provider for the message embed
func (m *MessageEmbed) SetProvider(provider *discordgo.MessageEmbedProvider) *MessageEmbed {
	m.embed.Provider = provider
	return m
}

// Sets Author for the message embed
func (m *MessageEmbed) SetAuthor(author *discordgo.MessageEmbedAuthor) *MessageEmbed {
	m.embed.Author = author
	return m
}

// // Sets Fields for the message embed
func (m *MessageEmbed) SetFields(fields []*discordgo.MessageEmbedField) *MessageEmbed {
	m.embed.Fields = fields
	return m
}

// Returns the discordgo.MessageEmbed struct
func (m *MessageEmbed) GetComponent() interface{} {
	return m.embed
}

// Returns the MessageEmbedAuthor struct
func NewEmbedAuthor(author string) *MessageEmbedAuthor {
	return &MessageEmbedAuthor{
		author: discordgo.MessageEmbedAuthor{
			Name: author,
		},
	}
}

// Sets URL for the MessageEmbedAuthor
func (a *MessageEmbedAuthor) SetURL(url string) *MessageEmbedAuthor {
	a.author.URL = url
	return a
}

// Sets Icon for the MessageEmbedAuthor
func (a *MessageEmbedAuthor) SetIcon(icon string) *MessageEmbedAuthor {
	a.author.IconURL = icon
	return a
}

//  Returns the discordgo.MessageEmbedAuthor
func (a *MessageEmbedAuthor) GetComponent() interface{} {
	return a.author
}

// Returns the Footer for the MessageEmbed
func NewEmbedFooter(text string, iconUrl string) *discordgo.MessageEmbedFooter {
	return &discordgo.MessageEmbedFooter{
		Text:    text,
		IconURL: iconUrl,
	}
}

// Returns the Image for the MessageEmbed
func NewEmbedImage(url string, width int, height int) *discordgo.MessageEmbedImage {
	if width == 0 || height == 0 {
		return &discordgo.MessageEmbedImage{
			URL: url,
		}
	}

	return &discordgo.MessageEmbedImage{
		URL:    url,
		Width:  width,
		Height: height,
	}
}

// Returns the Thumbnail for the MessageEmbed
func NewEmbedThumbnail(url string, width int, height int) *discordgo.MessageEmbedThumbnail {
	if width == 0 || height == 0 {
		return &discordgo.MessageEmbedThumbnail{
			URL: url,
		}
	}

	return &discordgo.MessageEmbedThumbnail{
		URL:    url,
		Width:  width,
		Height: height,
	}
}

// Returns the Video for the MessageEmbed
func NewEmbedVideo(url string, width int, height int) *discordgo.MessageEmbedVideo {
	if width == 0 || height == 0 {
		return &discordgo.MessageEmbedVideo{
			URL: url,
		}
	}

	return &discordgo.MessageEmbedVideo{
		URL:    url,
		Width:  width,
		Height: height,
	}
}

// Returns the Provider for the MessageEmbed
func NewEmbedProvider(url string, name string) *discordgo.MessageEmbedProvider {
	return &discordgo.MessageEmbedProvider{
		URL:  url,
		Name: name,
	}
}

// Returns the Field for the MessageEmbed
func NewEmbedFields(name string, value string, isInline bool) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: isInline,
	}
}
