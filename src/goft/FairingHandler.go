package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jhyehuang/crontab_server/src/result"
	"sync"
)

var fairingHandler *FairingHandler
var fairing_once sync.Once

func getFairingHandler() *FairingHandler {
	fairing_once.Do(func() {
		fairingHandler = &FairingHandler{}
	})
	return fairingHandler
}

type FairingHandler struct {
	fairings []Fairing
}

func NewFairingHandler() *FairingHandler {
	return &FairingHandler{}
}

func (this *FairingHandler) AddFairing(f ...Fairing) {
	if f != nil && len(f) > 0 {

		this.fairings = append(this.fairings, f...)
	}

}
func (this *FairingHandler) before(ctx *gin.Context) {
	for _, f := range this.fairings {
		err := f.OnRequest(ctx)
		if err != nil {
			Throw(err.Error(), 400, ctx)
		}
	}
}
func (this *FairingHandler) after(ctx *gin.Context, ret interface{}) interface{} {
	var result = ret
	for _, f := range this.fairings {
		r, err := f.OnResponse(result)
		if err != nil {
			Throw(err.Error(), 400, ctx)
		}
		result = r
	}
	return result
}

func (this *FairingHandler) handlerFairing(responder Responder, ctx *gin.Context) interface{} {
	this.before(ctx)
	var ret interface{}
	innerNode := getInnerRouter().getRoute(ctx.Request.Method, ctx.Request.URL.Path)
	var innerFairingHandler *FairingHandler
	if innerNode.fullPath != "" && innerNode.handlers != nil { //create inner fairinghandler for route-level middlerware.  hook like
		if fs, ok := innerNode.handlers.([]Fairing); ok {
			innerFairingHandler = NewFairingHandler()
			innerFairingHandler.AddFairing(fs...)
		}
	}
	// exec route-level middleware
	if innerFairingHandler != nil {
		innerFairingHandler.before(ctx)
	}
	if s1, ok := responder.(StringResponder); ok {
		ret = s1(ctx)
	}
	if s2, ok := responder.(JsonResponder); ok {
		ret = s2(ctx)
		if ret != nil {
			timeout := ret.(*result.Resp).Time
			if timeout > 0 {
				ctx.Writer.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d,public", timeout))
			}
		}
	}
	if s3, ok := responder.(VoidResponder); ok {
		s3(ctx)
		ret = struct{}{}
	}
	// exec route-level middleware
	if innerFairingHandler != nil {
		ret = innerFairingHandler.after(ctx, ret)
	}
	return getFairingHandler().after(ctx, ret)
}

// HandleFairing Deprecated ,please call FairingHandler.handlerFairing
func HandleFairing(responder Responder, ctx *gin.Context) interface{} {
	getFairingHandler().before(ctx)
	var ret interface{}
	if s1, ok := responder.(StringResponder); ok {
		ret = s1(ctx)
	}
	if s2, ok := responder.(JsonResponder); ok {
		ret = s2(ctx)
	}

	return getFairingHandler().after(ctx, ret)

}
