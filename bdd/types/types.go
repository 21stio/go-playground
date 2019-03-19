package types



type Language int32

const (
	DE Language = 0;
	EN Language = 1;
)

type FormattingKind int32

const (
	HTML FormattingKind = 0;
	PLAIN FormattingKind = 1;
	MARKDOWN FormattingKind = 2;
)

type ImageKind int32

const (
	ANY ImageKind = 0;
	AVATAR ImageKind = 1;
	HEADER ImageKind = 2;
	PROFILE ImageKind = 3;
	PREVIEW ImageKind = 4;
)

type LinkKind int32

const (
	DETAILS LinkKind = 0;
)

type Url struct {
	Value *string
	Hash *string`hash:"ignore"`
}

type Translation struct {
	Language *Language
	Formatting *FormattingKind
	Value *string
	Hash *string`hash:"ignore"`
}

type Link struct {
	Url *Url
	Kind *LinkKind
	Hash *string`hash:"ignore"`
}

type Image struct {
	Kind *ImageKind
	Url *Url
	Width *int32
	Height *int32
	Description *Text
	Hash *string`hash:"ignore"`
}

type Info struct {
	Name *Text
	Description *Text
	Images []Image
	Links []Link
	Hash *string`hash:"ignore"`
}

type Text struct {
	Language *Language
	Formatting *FormattingKind
	Value *string
	Translations []Translation
	Hash *string`hash:"ignore"`
}