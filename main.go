// File:		main.go
// Created by:	Hoven
// Created on:	2025-05-16
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-puzzles/puzzles/cores"
	"github.com/go-puzzles/puzzles/pflags"
	"github.com/go-puzzles/puzzles/pgin"
	"github.com/go-puzzles/puzzles/plog"

	httppuzzle "github.com/go-puzzles/puzzles/cores/puzzles/http-puzzle"
)

var (
	port = pflags.Int("port", 28080, "server run port")
)

func main() {
	pflags.Parse()

	engine := pgin.Default()

	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	engine.SetHTMLTemplate(tmpl)

	engine.Static("/static", "./static")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"CurrentTime": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	engine.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", map[string]interface{}{
			"CurrentTime": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	srv := cores.NewPuzzleCore(
		httppuzzle.WithCoreHttpPuzzle("/", engine),
	)
	plog.PanicError(cores.Start(srv, port()))
}
