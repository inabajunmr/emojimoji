package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"net/url"
	"os"

	"github.com/inabajunmr/gocolors"

	"github.com/inabajunmr/emosh"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type slackRequest struct {
	token       string
	teamID      string
	teamDomain  string
	channelID   string
	channelName string
	userID      string
	userName    string
	command     string
	text        string
	responseURL string
	triggerID   string
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// valification token
	token := os.Getenv("VERIFICATION_TOKEN")
	vals, _ := url.ParseQuery(request.Body)
	req := slackRequest{
		vals.Get("token"),
		vals.Get("team_id"),
		vals.Get("team_domain"),
		vals.Get("channel_id"),
		vals.Get("channel_name"),
		vals.Get("user_id"),
		vals.Get("user_name"),
		vals.Get("command"),
		vals.Get("text"),
		vals.Get("response_url"),
		vals.Get("trigger_id"),
	}

	if req.token != token {
		// invalid token
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Invalid token."),
			StatusCode: 200,
		}, nil
	}

	// TODO parse text
	// TODO bg,fg,text

	fcolorName := "Black"
	bcolorName := "White"
	fcolor, _ := gocolor.ValueOf(fcolorName, 255)
	bcolor, _ := gocolor.ValueOf(bcolorName, 255)
	// Generate Emoji
	emoji, _ := emosh.GenerateEmoji("test,", bcolor, fcolor)
	buf := &bytes.Buffer{}
	jpeg.Encode(buf, emoji, nil)

	// Slack api doc https://webapps.stackexchange.com/questions/89998/can-a-slackbot-create-emoji
	// TODO

	return events.APIGatewayProxyResponse{
		Body:       "Custom slash commands",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
