package handlers

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/helpers"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
)

func Verify(gctx *gin.Context) {
    log.Println()
    var email,_=gctx.Get("Email")
    var ctx,cancel=context.WithTimeout(context.Background(),time.Duration(15)*time.Minute)
    defer cancel()
    var verified,err=postgres.IsVerified(ctx,email.(string))
    if err!=nil{
        gctx.JSONP(http.StatusNotFound,"user does not exists")
        return
    }
    if verified{
        gctx.JSONP(http.StatusMethodNotAllowed,"user Already verified")
        return
    }
    jwt,err:=helpers.GenerateJWTVerify(email.(string))
    if err!=nil{
        gctx.JSONP(http.StatusInternalServerError,"couldnt make JWT")
        return
    }
    err=sendMail(jwt,email.(string))
    if err!=nil{
        return
    }
    gctx.JSONP(http.StatusOK,"mail sent")
}

func sendMail(jwt,mail string)error{
    var wg=&sync.WaitGroup{}
    wg.Add(1)
    var ch = make(chan error)
    helpers.SendMail(jwt,mail,wg,ch)
    wg.Wait()

    if err:=<-ch;err!=nil{
        return err
    }
    return nil
}
