package gopaper

// Return tag ID
func (t Tags) TagID() int {
	return t.ID
}

// Return tag name
func (t Tags) TagName() string {
	return t.Name
}

// Return alias of tag
func (t Tags) TagAlias() string {
	return t.Alias
}

// Return category id of a tag
func (t Tags) TagCategoryID() int {
	return t.CategoryID
}

// Return category name
func (t Tags) TagCategory() string {
	return t.Category
}

// Return tag content rating
func (t Tags) TagPurity() string {
	return t.Purity
}

// Return tag creation date
func (t Tags) TagDate() string {
	return t.Date
}
