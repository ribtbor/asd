// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}


type tCategory struct {}
var Category tCategory


func (_ tCategory) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Category.Index", args).URL
}

func (_ tCategory) Categories(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Category.Categories", args).URL
}


type tPornstar struct {}
var Pornstar tPornstar


func (_ tPornstar) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pornstar.Index", args).URL
}

func (_ tPornstar) Pornstars(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pornstar.Pornstars", args).URL
}


type tTag struct {}
var Tag tTag


func (_ tTag) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Tag.Index", args).URL
}

func (_ tTag) Tags(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Tag.Tags", args).URL
}

func (_ tTag) Page(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Tag.Page", args).URL
}


type tVideo struct {}
var Video tVideo


func (_ tVideo) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Video.Index", args).URL
}

func (_ tVideo) Videos(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Video.Videos", args).URL
}


