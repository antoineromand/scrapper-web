package common

import Models "scrapper-web/internal/model"

type HttpResponse struct {
    Error error
    Success bool
    Code int
    Message string
}

type HttpGetResponse struct {
    HttpResponse HttpResponse
    Data []Models.ScrapperOrder
}