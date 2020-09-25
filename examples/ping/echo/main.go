package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/philippgille/ln-paywall/ln"
	"github.com/philippgille/ln-paywall/storage"
	"github.com/philippgille/ln-paywall/wall"
)

func main() {
	e := ech