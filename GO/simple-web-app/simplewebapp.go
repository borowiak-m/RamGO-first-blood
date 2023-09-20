package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"log"
)

type ToDoList struct {
	ToDoCount int
	ToDos []string

}

func errorCheck