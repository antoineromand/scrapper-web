package dto

type ScrapperOrderDTO struct {
    Url string `json:"url"`
}

type ScrapperOrderSelectionDTO struct {
    Element   string  `json:"element"`
    Class     *string `json:"class"`
    Attribute *string `json:"attribute"`
    ElementId *string `json:"elementId"`
}
