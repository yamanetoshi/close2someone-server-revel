// app/controllers/api/v1/register

package controllers

import (
       "github.com/revel/revel"
//       "net/http"
)

type ApiV1Register struct {
     *revel.Controller
}

type ErrorResponse struct {
     Code    int    `json:"code"`
     Message string `json:"message"`
}

type Response struct {
     Results interface{} `json:"results"`
}

func (c ApiV1Register) Show(id int) revel.Result {
     r := Response{"Show"}
     return c.RenderJson(r)
}

func (c ApiV1Register) Post() revel.Result {
     r := Response{"Post"}
     return c.RenderJson(r)
}

func (c ApiV1Register) Delete(id int) revel.Result {
     r := Response{"Delete"}
     return c.RenderJson(r)
}
